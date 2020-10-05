package db

import (
	dbmodels "github.com/gangjun06/api.gangjun.dev/models/db"
	"github.com/gangjun06/api.gangjun.dev/utils"
	"gorm.io/gorm"
)

func GetInfo(key string) string {
	var result dbmodels.Info
	if err := utils.GetDB().Where("data_key = ?", key).Find(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			CreateInfo(key)
		}
		return ""
	}
	return result.Value
}

func CreateInfo(key string) error {
	return utils.GetDB().Create(&dbmodels.Info{DataKey: key, Value: ""}).Error
}

func SetInfo(key string, value string) error {
	return utils.GetDB().Model(&dbmodels.Info{}).Update("value = ", value).Error
}
