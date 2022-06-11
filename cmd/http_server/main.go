package main

import (
	"HttpServer/configs"
	"HttpServer/internal/app"
	"HttpServer/internal/server"
	"HttpServer/pkg/consts"
	"flag"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
)

var (
	configPath = flag.String("config", "conf/config.yaml", "config path")
)

func main() {
	configs.Init(*configPath)

	os.Setenv(consts.Version, "v1.0")

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
