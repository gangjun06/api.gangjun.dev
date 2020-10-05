package v1

import (
	c "github.com/gangjun06/api.gangjun.dev/controller/v1/info"
	"github.com/gin-gonic/gin"
)

func SetInfoRoutes(r *gin.RouterGroup) {
	r.GET("/discord", c.Discord)
	r.GET("/githubcontributionscount", c.GithubContributionsCount)
	r.GET("/value", c.GetValue)
}
