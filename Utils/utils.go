package Utils

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"net/http"
	"strings"
)

func UnmarshalJSON(c *gin.Context, data []byte, v interface{}, msg ...string) error {

	if len(msg) == 0 {
		msg = []string{"请求结构错误"}
	}
	//fmt.Println(string(data))
	err := json.Unmarshal(data, v)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"code": 7,
			"msg":  msg[0],
		})
		return err

	}
	return nil
}
func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func InSlice(s string, slice []string) bool {
	for _, v := range slice {
		if strings.Contains(s, v) {
			return true
		}
	}
	return false
}
func SearchSql(value map[string]interface{}, cut int) string {
	exclude := []string{"page", "pageSize"}
	//转义msg
	if value["msg"] != nil {
		msgstr := fmt.Sprintf("%v", value["msg"])
		encoded := base64.StdEncoding.EncodeToString([]byte(msgstr))
		value["msg"] = encoded
	}
	//NowTime := int(time.Now().Unix())
	var tempsql string
	for key, val := range value {

		valstr := fmt.Sprintf("%v", val)

		if !stringInSlice(key, exclude) && val != "" {
			tempsql = tempsql + key + " like '%" + valstr + "%'"
		}

		if !stringInSlice(key, exclude) && val != "" {
			tempsql = tempsql + " and "
		}
	}
	if tempsql != "" {
		tempsql = tempsql[:len(tempsql)-cut]
	} else {
		tempsql = ""
	}
	return tempsql
}
