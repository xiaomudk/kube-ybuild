package app

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"

	"github.com/xiaomudk/kube-ybuild/pkg/logs"
)

// App global app
type App struct {
	opts   options
	ctx    context.Context
	cancel func()
	mu     sync.Mutex
}

// New create a app globally
func New(opts ...Option) *App {
	o := options{
		ctx:    context.Background(),
		logger: logs.GetLogger(),
		// don not catch SIGKILL signal, need to waiting for kill self by other.
		sigs:            []os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT},
		registryTimeout: 10 * time.Second,
	}
	if id, err := uuid.NewUUID(); err == nil {
		o.id = id.String()
	}
	for _, opt := range opts {
		opt(&o)
	}

	ctx, cancel := context.WithCancel(o.ctx)
	return &App{
		opts:   o,
		ctx:    ctx,
		cancel: cancel,
	}
}

// Run start app
func (a *App) Run() error {

	eg, ctx := errgroup.WithContext(a.ctx)

	// start server
	wg := sync.WaitGroup{}
	for _, srv := range a.opts.servers {
		srv := srv
		eg.Go(func() error {
			// wait for stop signal
			<-ctx.Done()
			return srv.Stop(ctx)
		})
		wg.Add(1)
		eg.Go(func() error {
			wg.Done()
			return srv.Start(ctx)
		})
	}

	// watch signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, a.opts.sigs...)
	eg.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case s := <-quit:
				a.opts.logger.Infof("receive a quit signal: %s", s.String())
				err := a.Stop()
				if err != nil {
					a.opts.logger.Infof("failed to stop app, err: %s", err.Error())
					return err
				}
			}
		}
	})
	if err := eg.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}

	return nil
}

// Stop stops the application gracefully.
func (a *App) Stop() error {
	// cancel app
	if a.cancel != nil {
		a.cancel()
	}
	return nil
}
