package controllers

import (
	"HttpServer/internal/utils"
	pkgutils "HttpServer/pkg/utils"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
)

func QueryHandler(context *utils.Context) (*pkgutils.Response, error) {
	requestID := pkgutils.GetRequestID(context.Ctx)
	logrus.Infof("[%s] Start QueryHandler", requestID)
	request := context.Ctx.Request
	key := request.Header.Get("key")
	value, err := context.Server.DB.Get(key)

	if err != nil {
		logrus.Errorf("[%s] Query value of %s error: %s", requestID, key, err.Error())
		pkgutils.NewResponse("", fmt.Sprintf("%v", err), http.StatusBadRequest)
	}

	logrus.Infof("[%s] End QueryHandler", requestID)
	return pkgutils.NewResponse(value, "", http.StatusOK), nil
}
