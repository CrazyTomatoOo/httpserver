package router

import (
	inutils "HttpServer/internal/utils"
	"github.com/gin-gonic/gin"
)

const (
	apiPrefix  = "api"
	apiVersion = "v1"
)

type Router interface {
	NewGroup(string) *gin.RouterGroup
	Adaptor(inutils.Handler) gin.HandlerFunc
	RegisterAll()
}

type Group struct {
	group  *gin.RouterGroup
	router Router
}

func NewGroup(router Router, relativePath string) *Group {
	group := router.NewGroup(relativePath)
	return &Group{group: group, router: router}
}

func groupRegister(router Router) func(string) *Group {
	return func(relativePath string) *Group {
		return NewGroup(router, relativePath)
	}
}

func (g *Group) POST(relativePath string, handler inutils.Handler, middleware ...gin.HandlerFunc) {
	middleware = append(middleware, g.router.Adaptor(handler))
	g.group.POST(relativePath, middleware...)
}

func (g *Group) GET(relativePath string, handler inutils.Handler, middleware ...gin.HandlerFunc) {
	middleware = append(middleware, g.router.Adaptor(handler))
	g.group.GET(relativePath, middleware...)
}

func (g *Group) PUT(relativePath string, handler inutils.Handler, middleware ...gin.HandlerFunc) {
	middleware = append(middleware, g.router.Adaptor(handler))
	g.group.PUT(relativePath, middleware...)
}
