package router

import (
	"github.com/gin-gonic/gin"
	v1 "go-chains/api/v1"
	"go-chains/middleware"
)

func InitSystemRouter(Router *gin.RouterGroup) {
	SystemRouter := Router.Group("system").Use(middleware.OperationRecord())
	{
		SystemRouter.POST("getSystemConfig", v1.GetSystemConfig) // 获取配置文件内容
		SystemRouter.POST("setSystemConfig", v1.SetSystemConfig) // 设置配置文件内容
		SystemRouter.POST("getServerInfo", v1.GetServerInfo)     // 获取服务器信息
	}
}
