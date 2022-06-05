package utils

import (
	"HttpServer/pkg/consts"
	"fmt"

	"github.com/gin-gonic/gin"
)

// Response ..
type Response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
}

func NewResponse(data interface{}, msg string, code int) *Response {
	return &Response{
		Data:    data,
		Message: msg,
		Code:    code,
	}
}

func WriteResponse(context *gin.Context, res *Response, msg error) {
	requestId := GetRequestID(context)
	context.Set(consts.XRequestID, requestId)

	var code int
	if msg != nil && res.Code == 0 {
		code = 500
		res.Message = fmt.Sprintf("%v", msg)
	} else {
		if res.Code != 0 {
			code = res.Code
			res.Code *= 100
			res.Message = fmt.Sprintf("%v", msg)
		} else {
			code = 200
		}
	}

	context.JSON(code, res)
}

func GetRequestID(context *gin.Context) string {
	v, exist := context.Get(consts.XRequestID)
	if !exist {
		return ""
	}

	if _, ok := v.(string); !ok {
		return ""
	}

	return v.(string)
}
