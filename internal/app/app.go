package app

import (
	"HttpServer/configs"
	"HttpServer/internal/router"
	"HttpServer/internal/server"
	"HttpServer/pkg/metrics"
	"HttpServer/pkg/middleware"
	"context"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type App struct {
	http   *http.Server
	engine *gin.Engine
}

func NewApp() *App {
	return &App{
		engine: gin.New(),
	}
}

func (app *App) Run(s *server.Server) {
	app.engine.Use(gin.Logger(), gin.Recovery())
	app.engine.Use(middleware.GenRequestID())
	app.engine.Use(middleware.AddVersion())
	app.engine.Use(middleware.AccessLog())

	router.RegisterAll(app.engine, s)

	app.engine.GET("/metrics", metrics.PromHandler(promhttp.Handler()))

	app.http = &http.Server{
		Addr:    viper.GetString(configs.Addr),
		Handler: app.engine,
	}

	if err := app.http.ListenAndServe(); err != nil {
		if err != http.ErrServerClosed {
			logrus.Error(err)
			os.Exit(-1)
		}
		logrus.Info("server closed under request")
	}
}

func (app *App) Close() {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		3*time.Second,
	)
	defer cancel()

	if err := app.http.Shutdown(ctx); err != nil {
		logrus.Infof("shutdown http server err %v", err)
	}
}
