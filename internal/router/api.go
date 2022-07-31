package router

import (
	"HttpServer/internal/controllers"
	"HttpServer/internal/server"
	"HttpServer/internal/utils"
	pkgutils "HttpServer/pkg/utils"
	"github.com/gin-gonic/gin"
)

type APIRouter struct {
	engine  *gin.Engine
	prefix  string
	version string
	group   func(string) *Group
	server  *server.Server
}

type Handler func(context *utils.Context) (*pkgutils.Response, error)

func NewAPIRouter(engine *gin.Engine, server *server.Server) *APIRouter {
	apiRouter := &APIRouter{
		engine:  engine,
		prefix:  apiPrefix,
		version: apiVersion,
		server:  server,
	}
	apiRouter.group = groupRegister(apiRouter)

	return apiRouter
}

func (api *APIRouter) NewGroup(relativePath string) *gin.RouterGroup {
	group := api.engine.Group(api.prefix).Group(api.version)
	return group.Group(relativePath)
}

func (api *APIRouter) Adaptor(handler utils.Handler) gin.HandlerFunc {
	return func(context *gin.Context) {
		data, err := handler(&utils.Context{
			Ctx:    context,
			Server: api.server,
		})
		pkgutils.WriteResponse(context, data, err)
	}
}

func (api *APIRouter) RegisterAll() {
	apiGroup := api.group("zsj")
	apiGroup.GET("/healthz", utils.ViewHandler(controllers.HealthzHandler))
	apiGroup.GET("/data", utils.ViewHandler(controllers.QueryHandler))
	apiGroup.POST("/data", utils.ViewHandler(controllers.InsertHandler))
	
}
