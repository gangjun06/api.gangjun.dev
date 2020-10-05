package info

import (
	"github.com/gangjun06/api.gangjun.dev/db"
	"github.com/gangjun06/api.gangjun.dev/utils"
	resutils "github.com/gangjun06/api.gangjun.dev/utils/res"
	"github.com/gin-gonic/gin"
	"github.com/imroc/req"
)

func Discord(c *gin.Context) {
	r := resutils.New(c)

	discordID := db.GetInfo("discord_id")
	discordConfig := utils.GetConfig().Discord
	resp, err := req.Get("https://discord.com/api/v8/users/"+discordID, req.Header{
		"Authorization": "Bot " + discordConfig.Bot,
	})
	if err != nil {
		r.SendError(resutils.ERR_SERVER, "Error while get information")
		return
	}

	var data map[string]interface{}
	if err := resp.ToJSON(&data); err != nil {
		r.SendError(resutils.ERR_SERVER, "Error while parse json")
		return
	}
	r.Response(data)
}
