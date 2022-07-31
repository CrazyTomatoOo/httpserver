package main

import (
	"HttpServer/configs"
	"HttpServer/internal/app"
	"HttpServer/internal/server"
	"HttpServer/pkg/metrics"
	"flag"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
)

var (
	configPath = flag.String("config", "conf/config.yaml", "config path")
)

func main() {
	flag.Parse()
	configs.Init(*configPath)

	level, err := log.ParseLevel(viper.GetString("log_level"))
	if err != nil {
		log.Error("Unknown log level, set info")
		level = log.InfoLevel
	}
	log.SetLevel(level)

	metrics.Register()

	s := server.NewServer()
	s.Init()

	App := app.NewApp()
	go App.Run(s)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	App.Close()
	log.Error("http server closed")
}
