package v1

import (
	c "github.com/gangjun06/api.gangjun.dev/controller/v1/utils"
	"github.com/gin-gonic/gin"
)

func SetUtilsRoutes(r *gin.RouterGroup) {
	r.POST("/sendemailtome", c.SendEmailToMe)
}
