package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/gangjun06/api.gangjun.dev/middlewares"
	"github.com/gangjun06/api.gangjun.dev/models"
	"github.com/gangjun06/api.gangjun.dev/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	etcInit()
	applyConfig()
	startServer()
}

func etcInit() {
	rand.Seed(time.Now().Unix())
}

func applyConfig() {
	rawConfig, err := ioutil.ReadFile("config.toml")
	if err != nil {
		log.Fatalln("Failed to load config.")
	}

	var config models.Config
	if _, err := toml.Decode(string(rawConfig), &config); err != nil {
		log.Fatalln("Failed to parsing config.")
	}
	utils.SetConfig(&config)
}

func startServer() {
	config := utils.GetConfig().Server
	if config.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)

	}
	r := gin.Default()
	r.Use(middlewares.Cors())
	r.Run(":" + strconv.Itoa(config.Port))
}
