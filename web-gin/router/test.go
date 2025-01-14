package router

import (
	"github.com/gin-gonic/gin"
	"go-micro-frame-web/api/ceshi"
	"go-micro-frame-web/middlewares"
)

func InitTestRouter(Router *gin.RouterGroup){
	UserRouter := Router.Group("test").Use(middlewares.Trace())
	{
		UserRouter.GET("send-mq", ceshi.SendMq)
	}
}