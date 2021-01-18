package service

import (
	"errors"
	"go-chains/chains/models"
	"go-chains/global"
	"gorm.io/gorm"
)

func CreateMetaData(meta *models.ChainMetaData) (err error) {
	queryResult := global.GVA_DB.Where("chain_id = ? AND chain_name = ? ",
		meta.ChainId, meta.ChainName).First(&models.ChainMetaData{})

	if !errors.Is(queryResult.Error, gorm.ErrRecordNotFound) {
		return errors.New("duplicated chain Id existing ")
	}
	return global.GVA_DB.Create(&meta).Error
}

func QueryAllChainMetaData() []models.ChainMetaData {
	result := make([]models.ChainMetaData, 0)
	global.GVA_DB.Find(&result)
	return result
}

func UpdateChainMetaData(meta *models.ChainMetaData) error {
	result := &models.ChainMetaData{}
	queryResult := global.GVA_DB.Where("chain_id = ? AND chain_name = ?",
		meta.ChainId, meta.ChainName).First(result)

	if !errors.Is(queryResult.Error, gorm.ErrRecordNotFound) {
		meta.ID = result.ID
		global.GVA_DB.Updates(meta)
	}
	return global.GVA_DB.Create(&meta).Error
}
