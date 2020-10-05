package info

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	resutils "github.com/gangjun06/api.gangjun.dev/utils/res"
	"github.com/gin-gonic/gin"
	"github.com/imroc/req"
)

func GithubContributionsCount(c *gin.Context) {
	r := resutils.New(c)

	resp, err := req.Get("https://proxy.gangjun.dev/github/users/gangjun06/contributions")
	if err != nil {
		r.SendError(resutils.ERR_SERVER, err.Error())
		return
	}

	doc, err := goquery.NewDocumentFromResponse(resp.Response())
	if err != nil {
		r.SendError(resutils.ERR_SERVER, err.Error())
		return
	}

	text := strings.Split(strings.TrimSpace(doc.Find("h2").Text()), " ")[0]
	result, _ := strconv.Atoi(text)
	r.Response(map[string]interface{}{"value": result})
}
