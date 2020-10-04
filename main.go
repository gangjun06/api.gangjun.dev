package main

import (
	"math/rand"
	"time"

	"github.com/gangjun06/api.gangjun.dev/middlewares"
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
}

func startServer() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.Use(middlewares.Cors())
	r.Run(":9096")
}
