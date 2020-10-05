package utils

import (
	reqmodels "github.com/gangjun06/api.gangjun.dev/models/req"
	"github.com/gangjun06/api.gangjun.dev/utils"
	resutil "github.com/gangjun06/api.gangjun.dev/utils/res"
	"github.com/gin-gonic/gin"
	"github.com/imroc/req"
)

func SendEmailToMe(c *gin.Context) {
	body := c.MustGet("body").(*reqmodels.SendEmailToMe)
	r := resutil.New(c)

	captchaSecretKey := utils.GetConfig().ReCAPTCHA.SecretKey

	header := req.Header{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	param := req.Param{
		"secret":   captchaSecretKey,
		"response": body.Captcha,
	}

	resp, err := req.Post("ttps://www.google.com/recaptcha/api/siteverify", param, header)
	if err != nil {
		r.SendError(resutil.ERR_BAD_REQUEST, "error wlhile request to recaptcha")
	}

	var data map[string]interface{}
	resp.ToJSON(data)
}
