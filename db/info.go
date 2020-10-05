package db

import (
	dbmodels "github.com/gangjun06/api.gangjun.dev/models/db"
	"github.com/gangjun06/api.gangjun.dev/utils"
)

func GetInfo(key string) string {
	var result dbmodels.Info
	if err := utils.GetDB().Where("key = ?", key).Find(&result).Error; err != nil {
		CreateInfo(key)
		return ""
	}
	return result.Value
}

func CreateInfo(key string) error {
	return utils.GetDB().Create(&dbmodels.Info{Key: key, Value: ""}).Error
}

func SetInfo(key string, value string) error {
	return utils.GetDB().Model(&dbmodels.Info{}).Update("value = ", value).Error
}
