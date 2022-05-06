package app

import (
	"context"
	"net/url"
	"os"
	"time"

	"github.com/xiaomudk/kube-ybuild/pkg/logs"
	"github.com/xiaomudk/kube-ybuild/pkg/transport"
)

// Option is func for application
type Option func(o *options)

// options is an application options
type options struct {
	id        string
	name      string
	version   string
	metadata  map[string]string
	endpoints []*url.URL

	sigs []os.Signal
	ctx  context.Context

	logger logs.Logger

	registryTimeout time.Duration
	servers         []transport.Server
}

// WithID with app id
func WithID(id string) Option {
	return func(o *options) {
		o.id = id
	}
}

// WithName .
func WithName(name string) Option {
	return func(o *options) {
		o.name = name
	}
}

// WithVersion with a version
func WithVersion(version string) Option {
	return func(o *options) {
		o.version = version
	}
}

// WithContext with a context
func WithContext(ctx context.Context) Option {
	return func(o *options) {
		o.ctx = ctx
	}
}

// WithSignal with some system signal
func WithSignal(sigs ...os.Signal) Option {
	return func(o *options) {
		o.sigs = sigs
	}
}

// WithMetadata with service metadata.
func WithMetadata(md map[string]string) Option {
	return func(o *options) { o.metadata = md }
}

// WithEndpoint with service endpoint.
func WithEndpoint(endpoints ...*url.URL) Option {
	return func(o *options) { o.endpoints = endpoints }
}

// WithLogger .
func WithLogger(logger logs.Logger) Option {
	return func(o *options) {
		o.logger = logger
	}
}

// WithServer with a server , http or grpc
func WithServer(srv ...transport.Server) Option {
	return func(o *options) {
		o.servers = srv
	}
}
