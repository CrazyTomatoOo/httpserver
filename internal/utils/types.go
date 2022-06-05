package utils

import (
	"HttpServer/internal/server"
	"HttpServer/pkg/consts"
	"HttpServer/pkg/utils"
	"database/sql/driver"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type Context struct {
	Ctx    *gin.Context
	Server *server.Server
}

type Handler func(context *Context) (*utils.Response, error)

func ViewHandler(f func(context *Context) (*utils.Response, error)) Handler {
	return f
}

type Time time.Time

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+consts.TimeFormat+`"`, string(data), time.Local)
	*t = Time(now)

	return
}

func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(consts.TimeFormat)+2)
	b = append(b, '"')
	b = t.local().AppendFormat(b, consts.TimeFormat)
	b = append(b, '"')

	return b, nil
}

func (t Time) String() string {
	return t.local().Format(consts.TimeFormat)
}

func (t Time) Format(format string) string {
	return t.local().Format(format)
}

func (t Time) local() time.Time {
	local, _ := time.LoadLocation(consts.Zone)
	return time.Time(t).In(local)
}

// Value ...
func (t Time) Value() (driver.Value, error) {
	var zeroTime time.Time
	var ti = time.Time(t)

	if ti.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}

	return ti, nil
}

// Scan value time.Time 注意是指针类型 method
func (t *Time) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = Time(value)
		return nil
	}

	return fmt.Errorf("can not convert %v to timestamp", v)
}

