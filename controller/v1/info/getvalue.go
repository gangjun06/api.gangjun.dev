package info

import (
	"fmt"
	"strings"

	"github.com/gangjun06/api.gangjun.dev/db"

	resutil "github.com/gangjun06/api.gangjun.dev/utils/res"
	"github.com/gin-gonic/gin"
)

func GetValue(c *gin.Context) {
	r := resutil.New(c)

	value := c.Query("key")
	fmt.Println(value)

	list := strings.Split(value, ",")
	listlen := len(list)
	if listlen < 1 || listlen > 10 {
		r.SendError(resutil.ERR_BAD_REQUEST, "")
		return
	}

	result := map[string]interface{}{}

	for _, item := range list {
		value := db.GetInfo(strings.TrimSpace(item))
		result[item] = value
	}

	r.Response(result)
}
