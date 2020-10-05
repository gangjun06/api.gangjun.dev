package middlewares

import (
	"github.com/gangjun06/api.gangjun.dev/utils/res"
	"github.com/gin-gonic/gin"
)

func VerifyRequest(data interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		r := res.New(c)
		if err := c.ShouldBindJSON(data); err != nil {
			r.SendError(res.ERR_BAD_REQUEST, err.Error())
			return
		}
		c.Set("body", data)
	}
}
