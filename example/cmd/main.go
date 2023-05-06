package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-leo/errors"
	"github.com/go-leo/errors/example/api"
	"github.com/go-leo/errors/example/code"
	"github.com/go-leo/errors/example/internal/service"
	"github.com/go-leo/leo/global"
	"github.com/go-leo/leo/log"
	middlewarelog "github.com/go-leo/leo/middleware/log"
)

func main() {
	r := gin.Default()
	r.Use(
		middlewarelog.GinMiddleware(
			log.FromContextOrDiscard, middlewarelog.WithPayloadWhenError(),
		),
	)
	r.GET("/user", GetUser)
	_ = r.Run()
}

func GetUser(c *gin.Context) {
	r := new(api.GetUserReq)
	if err := c.BindQuery(r); err != nil {
		err := errors.WithCode(err, code.ErrBind)
		_ = c.Error(err) // 设置后会有统一日志输出
		coder := errors.ParseCoder(err)
		c.AbortWithStatusJSON(coder.HTTPStatus(), coder)
		global.Logger().Errorf("bind query failed: %-v", err)
		return
	}
	svc := service.NewUserSvc()
	reply, err := svc.GetUser(r)
	if err != nil {
		_ = c.Error(err)
		coder := errors.ParseCoder(err)
		c.JSON(coder.HTTPStatus(), coder)
		global.Logger().Errorf("GetUser failed: %-v", err)
		return
	}
	c.JSON(200, reply)
}
