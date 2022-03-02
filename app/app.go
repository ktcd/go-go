package app

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/ktcd/go-go/app/config"
	"github.com/ktcd/go-go/app/logger"
	"github.com/ktcd/go-go/app/meta"
	"github.com/ktcd/go-go/app/router"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type App struct {
	meta   *meta.Meta
	logger *logrus.Logger
	router *chi.Mux
}

func Init() *App {
	config.Init()

	var (
		meta   = meta.Init()
		logger = logger.Init()
		router = router.Init()
		app    = &App{
			meta:   meta,
			logger: logger,
			router: router,
		}
	)

	return app
}

func (app *App) GetEnvironment() string {
	return viper.GetString("APP_ENV")
}

func (app *App) GetVersion() string {
	var (
		env     = app.GetEnvironment()
		version = viper.GetString("APP_VERSION")
	)

	if env != "production" {
		version += "-" + env
	}

	return version
}

func (app *App) GetMeta() *meta.Meta {
	return app.meta
}

func (app *App) GetHttpRouter() *chi.Mux {
	return app.router
}

func (app *App) GetLogger() *logrus.Logger {
	return app.logger
}

func (app *App) Run() error {
	var port = viper.GetString("APP_PORT")
	var host = "0.0.0.0:" + port
	return http.ListenAndServe(host, app.router)
}
