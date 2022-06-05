package router

import (
	"HttpServer/internal/server"
	"github.com/gin-gonic/gin"
)

func RegisterAll(engine *gin.Engine, server *server.Server) {
	routers := []Router{
		NewAPIRouter(engine, server),
	}
	for _, route := range routers {
		route.RegisterAll()
	}
}
