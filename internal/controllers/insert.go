package controllers

import (
	"HttpServer/internal/utils"
	pkgutils "HttpServer/pkg/utils"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func InsertHandler(context *utils.Context) (*pkgutils.Response, error) {
	requestID := pkgutils.GetRequestID(context.Ctx)
	logrus.Infof("[%s] Start InsertHandler", requestID)

	m := map[string]string{}
	body, err := ioutil.ReadAll(context.Ctx.Request.Body)
	if err != nil {
		logrus.Errorf("[%s] read request body err: %s", requestID, err.Error())
		return pkgutils.NewResponse("", err.Error(), http.StatusInternalServerError), err
	}

	err = json.Unmarshal(body, &m)
	if err != nil {
		logrus.Errorf("[%s] Unmarshal body err: %s", requestID, err.Error())
		return pkgutils.NewResponse("", err.Error(), http.StatusBadRequest), err
	}

	err = context.Server.DB.BatchInsert(m)
	if err != nil {
		logrus.Errorf("[%s] Insert db error: %s", requestID, err.Error())
		return pkgutils.NewResponse("", err.Error(), http.StatusInternalServerError), err
	}

	logrus.Infof("[%s] End InsertHandler", requestID)
	return pkgutils.NewResponse("", "", http.StatusOK), nil
}
