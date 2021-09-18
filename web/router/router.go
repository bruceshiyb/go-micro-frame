package router

import (
	"github.com/gin-gonic/gin"
	"go-micro-frame-web/api/user"
	"go-micro-frame-web/middlewares"
)

func InitUserRouter(Router *gin.RouterGroup){
	UserRouter := Router.Group("user").Use(middlewares.Trace())
	{
		UserRouter.GET("list", user.GetUserList)
	}
}