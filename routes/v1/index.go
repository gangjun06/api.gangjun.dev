package v1

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes(g *gin.RouterGroup) {
	infoRoutes := g.Group("info")
	SetInfoRoutes(infoRoutes)

	utilsRoutes := g.Group("utils")
	SetUtilsRoutes(utilsRoutes)
}
