package middleware

import (
	"HttpServer/configs"
	"HttpServer/pkg/consts"
	"HttpServer/pkg/utils"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"time"
)

func GenRequestID() gin.HandlerFunc {
	return func(context *gin.Context) {
		uid := uuid.NewV4().String()
		context.Set(consts.XRequestID, uid)
		context.Next()
	}
}

func AddVersion() gin.HandlerFunc {
	return func(context *gin.Context) {
		version := viper.GetString(configs.Version)
		if v := os.Getenv(consts.Version); v != consts.EmptyString {
			version = v
		}

		context.Writer.Header().Set(consts.Version, version)
		context.Next()
	}
}

func AccessLog() gin.HandlerFunc {
	return func(context *gin.Context) {
		local, _ := time.LoadLocation(consts.Zone)

		logrus.Infof("%s %s %s %s",
			time.Now().In(local).Format(consts.TimeFormat),
			utils.GetClientIP(context.Request),
			context.Request.Method,
			context.Request.URL,
		)

		context.Next()
	}
}
