package utils

import (
	"github.com/gangjun06/api.gangjun.dev/models"
)

var (
	gConfig *models.Config
)

func SetConfig(config *models.Config) {
	gConfig = config
}

func GetConfig() *models.Config {
	return gConfig
}
