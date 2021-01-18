package router

import (
	"github.com/gin-gonic/gin"
	v1 "go-chains/api/v1"
	"go-chains/middleware"
)

func InitJwtRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("jwt").Use(middleware.OperationRecord())
	{
		ApiRouter.POST("jsonInBlacklist", v1.JsonInBlacklist) // jwt加入黑名单
	}
}
