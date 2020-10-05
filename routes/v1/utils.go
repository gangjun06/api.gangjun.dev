package v1

import (
	c "github.com/gangjun06/api.gangjun.dev/controller/v1/utils"
	m "github.com/gangjun06/api.gangjun.dev/middlewares"
	"github.com/gangjun06/api.gangjun.dev/models/req"
	"github.com/gin-gonic/gin"
)

func SetUtilsRoutes(r *gin.RouterGroup) {
	r.POST("/sendemailtome", m.VerifyRequest(&req.SendEmailToMe{}), c.SendEmailToMe)
}
