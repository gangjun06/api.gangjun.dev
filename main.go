package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strconv"
	"time"

	dbmodels "github.com/gangjun06/api.gangjun.dev/models/db"
	v1 "github.com/gangjun06/api.gangjun.dev/routes/v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/BurntSushi/toml"
	"github.com/gangjun06/api.gangjun.dev/middlewares"
	"github.com/gangjun06/api.gangjun.dev/models"
	"github.com/gangjun06/api.gangjun.dev/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	etcInit()
	applyConfig()
	initDB()
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
	version1 := r.Group("/v1")
	v1.InitRoutes(version1)
	r.Run(":" + strconv.Itoa(config.Port))
}

func initDB() {
	log.Println("Initializing Database...")
	dbConfig := utils.GetConfig().DB
	connectionInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", dbConfig.Username, dbConfig.Password, dbConfig.Hostname, dbConfig.Port, dbConfig.DBName)
	db, err := gorm.Open(mysql.Open(connectionInfo), &gorm.Config{})
	if err != nil {
		log.Fatalln("Failed to open database.")
	}
	utils.SetDB(db)
	log.Print("Successfully Connected To Database")

	var models = []interface{}{&dbmodels.Info{}}

	if err := db.AutoMigrate(models...); err != nil {
		log.Fatalln("Failed to perform AutoMigrate.")
	}
	log.Print("Successfully performed AutoMigrate")
}
