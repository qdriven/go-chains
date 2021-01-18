package router

import (
	"github.com/gin-gonic/gin"
	v1 "go-chains/api/v1"
	"go-chains/middleware"
)

func InitEmailRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("email").Use(middleware.OperationRecord())
	{
		UserRouter.POST("emailTest", v1.EmailTest) // 发送测试邮件
	}
}
