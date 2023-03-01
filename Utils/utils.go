package Utils

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"net/http"
	"strconv"
	"time"
)

func UnmarshalJSON(c *gin.Context, data []byte, v interface{}, msg ...string) error {

	if len(msg) == 0 {
		msg = []string{"結構錯誤"}
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
func SearchSql(value map[string]interface{}, cut int) string {
	resouce := []string{"primogem", "acquaint", "intertwined", "level"}
	exclude := []string{"page", "pageSize", "primogemmax", "acquaintmax", "intertwinedmax", "levelmax"}
	//转义msg
	if value["msg"] != nil {
		msgstr := fmt.Sprintf("%v", value["msg"])
		encoded := base64.StdEncoding.EncodeToString([]byte(msgstr))
		value["msg"] = encoded
	}
	NowTime := int(time.Now().Unix())
	var tempsql string
	for key, val := range value {

		valstr := fmt.Sprintf("%v", val)
		if stringInSlice(key, resouce) && val != "" {
			max := value[key+"max"]
			if max == "" {
				max = "2100000000"
			}
			maxstr := fmt.Sprintf("%v", max)
			tempsql = tempsql + key + ">=" + valstr + " and " + key + "<=" + maxstr
		} else if key == "runtime" && val != "" {
			valformat, _ := strconv.Atoi(valstr)
			tempsql = tempsql + "(" + strconv.Itoa(int(time.Now().Unix())) + "- logintime)>=" + strconv.Itoa(valformat*3600)
		} else if key == "lastupdate" && val != "" {
			valformat, _ := strconv.Atoi(valstr)
			tempsql = tempsql + "time <= " + strconv.Itoa(valformat/1000+86399)
		} else if key == "machine" && val != "" {
			tempsql = tempsql + "hwid in (select hwid from machine where name='" + valstr + "')"
		} else if key == "online" && len(valstr) > 0 {
			//fmt.Println("online")
			if valstr == "0" {
				tempsql = tempsql + "time+60>" + strconv.Itoa(NowTime)
			} else if valstr == "1" {
				tempsql = tempsql + "time+60<" + strconv.Itoa(NowTime)
			}
		} else if !stringInSlice(key, exclude) && val != "" {
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
