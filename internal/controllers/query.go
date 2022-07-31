package controllers

import (
	"HttpServer/internal/utils"
	"HttpServer/pkg/metrics"
	pkgutils "HttpServer/pkg/utils"
	"fmt"
	"github.com/sirupsen/logrus"
	"math/rand"
	"net/http"
	"time"
)

func QueryHandler(context *utils.Context) (*pkgutils.Response, error) {
	requestID := pkgutils.GetRequestID(context.Ctx)
	logrus.Infof("[%s] Start QueryHandler", requestID)

	timer := metrics.NewTimer()
	defer timer.ObserveTotal()

	delay := rand.Intn(1990) + 10
	time.Sleep(time.Millisecond * time.Duration(delay))

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
