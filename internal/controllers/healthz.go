package controllers

import (
	"HttpServer/internal/utils"
	pkgutils "HttpServer/pkg/utils"
	"github.com/sirupsen/logrus"
	"net/http"
)

func HealthzHandler(context *utils.Context) (*pkgutils.Response, error) {
	logrus.Infof("[%s] App OK.", pkgutils.GetRequestID(context.Ctx))
	return pkgutils.NewResponse("", "App OK!!!!", http.StatusOK), nil
}
