package Middleware

import (
	"face/Utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		token := c.Request.Header.Get("X-TOKEN")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Headers", "*")

			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			//c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			//c.Header("Access-Control-Expose-Headers", "*")
			c.Header("Access-Control-Allow-Credentials", "false")
			c.Set("content-type", "application/json")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		exclude := []string{"Heart", "Detected", "Login", "Captcha", "GetSignInfo"}
		if method != "OPTIONS" && Utils.InSlice(c.Request.URL.String(), exclude) == false {
			//fmt.Println(token)
			resp := Utils.Decode_jwt_token(token, "username")
			//fmt.Println(resp)
			if resp == "106" || resp == "108" {
				c.Abort()
				c.JSON(403, gin.H{
					"msg": "Token过期",
				})
				return
			}

		}

		c.Next()
	}
}
