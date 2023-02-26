package Middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		//token := c.Request.Header.Get("x-token")
		//獲得密碼

		//var Setting Models.Setting
		//db := Global.DB

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

		//api.GET("/getcoremd5.php", Controller.GetCoreMd5)       //获取中控CoreMd5
		//api.GET("/Global_Config.json", Controller.GetConfig)    //获取config.json
		//api.GET("/getdeviceinfo.php", Controller.GetDeviceInfo) //取号接口
		//api.POST("/Heart", Controller.Heart)                    //心跳接口
		//api.POST("/Update", Controller.Update)                  //Update更新信息接口
		//api.POST("/GetScript", Controller.GetScript)            //拉脚本接口
		//api.POST("/Process", Controller.Process)                //进度Get/Set接口
		//exclude := []string{"getcoremd5", "Global_Config", "getdeviceinfo", "Heart", "Update", "GetScript", "Process", "res"}
		//exclude := []string{"InitDb", "getcoremd5", "Global_Config", "getdeviceinfo", "heart", "update", "getscript", "process", "res"}
		//pathexclude := []string{"api", "res"}
		////fmt.Println(c.Request.URL.Path)
		////fmt.Println(c.Request.URL.String())
		////fmt.Println(Utils.InSlice(c.Request.URL.Path, pathexclude))
		//
		//if method != "OPTIONS" && Utils.InSlice(c.Request.URL.String(), exclude) == false && Utils.InSlice(c.Request.URL.Path, pathexclude) {
		//	//fmt.Println("触发鉴权")
		//	db.Where("name = ?", "remotepassword").Take(&Setting)
		//	Pwd := Setting.Value
		//	if token != Pwd {
		//		c.Abort()
		//		c.JSON(200, gin.H{
		//			//"msg": "当前访问路由" + c.Request.URL.String() + " 輸入的密碼不正確 輸入的密碼是:" + token,
		//			"msg": "輸入的密碼不正確 輸入的密碼是:" + token,
		//		})
		//	}
		//
		//}

		c.Next()
	}
}
