package utils

import (
	"fmt"

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

	resp, err := req.Post("https://www.google.com/recaptcha/api/siteverify", param, header)
	if err != nil {
		r.SendError(resutil.ERR_SERVER, "error wlhile request to recaptcha")
		return
	}

	var data map[string]interface{}
	resp.ToJSON(&data)
	if !data["success"].(bool) {
		r.SendError(resutil.ERR_BAD_REQUEST, "captcha token is invalid")
		return
	}

	if err := utils.SendEmailToMe(body.Title, "Send By: "+body.Email+"\n\n"+body.Text); err != nil {
		fmt.Println(err.Error())
		r.SendError(resutil.ERR_SERVER, "Error sending email")
		return
	}

	r.Response(map[string]interface{}{})
}
