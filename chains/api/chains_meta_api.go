package api

import (
	"github.com/gin-gonic/gin"
	"go-chains/chains/api/validator"
	"go-chains/chains/models"
	"go-chains/chains/service"
	"go-chains/global"
	"go-chains/model/response"
	"go-chains/utils"
	"go.uber.org/zap"
)

// @Tags ChainMetaDataApi
// @Summary 创建Chain的基础信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body models.ChainMetaData true "	ChainName,ChainI,ChainType,RPCUrl string"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /chains/chainMetaData [post]
func CreateChainMetaData(c *gin.Context) {
	var request = &models.ChainMetaData{}
	_ = c.ShouldBindJSON(&request)
	if err := utils.Verify(request, validator.ChainMetaValidator); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := service.CreateMetaData(request); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags ChainMetaDataApi
// @Summary 获取Chain MetaData信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /chains/chainMetaData [get]
func QueryChainMetaData(c *gin.Context) {
	result := service.QueryAllChainMetaData()
	response.OkWithData(result, c)
}
