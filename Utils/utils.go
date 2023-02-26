package Utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"net/http"
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
