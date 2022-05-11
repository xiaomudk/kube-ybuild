package server

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/spf13/pflag"

	"github.com/xiaomudk/kube-ybuild/internal/repository"
	"github.com/xiaomudk/kube-ybuild/internal/service"
	logger "github.com/xiaomudk/kube-ybuild/pkg/logs"

	"github.com/xiaomudk/kube-ybuild/internal/config"
	"github.com/xiaomudk/kube-ybuild/internal/middleware"
	"github.com/xiaomudk/kube-ybuild/internal/model"
	"github.com/xiaomudk/kube-ybuild/internal/routers"
	"github.com/xiaomudk/kube-ybuild/pkg/app"
	"github.com/xiaomudk/kube-ybuild/pkg/transport/http"
)

var (
	conf    = pflag.StringP("config dir", "c", "config/kube-ybuild.yaml", "config path.")
	env     = pflag.StringP("env name", "e", "", "env var name.")
	version = pflag.BoolP("version", "v", false, "show version info.")
)

type APIValidator struct {
	validator *validator.Validate
}

func (apiVal *APIValidator) Validate(i interface{}) error {
	return apiVal.validator.Struct(i)
}

// NewHTTPServer creates an HTTP server
func NewHTTPServer() *app.App {
	e := echo.New()
	config.Init(*conf)
	e.Logger = logger.NewEchoLogger(config.Conf.Log)
	e.Validator = &APIValidator{validator: validator.New()}

	// -------------- init resource -------------
	routers.Init(e)
	middleware.Init(e)
	// init db
	model.Init()
	service.Svc = service.New(repository.New(model.GetDB()))
	model.MigrateDatabase(model.GetDB())

	srv := http.NewServer(
		http.WithAddress(config.Conf.Address),
		http.WithReadTimeout(10*time.Second),
		http.WithWriteTimeout(10*time.Second),
	)

	srv.Handler = e

	// start app
	newApp := app.New(
		app.WithLogger(e.Logger),
		app.WithServer(
			// init http server
			srv,
		),
	)
	return newApp

}
