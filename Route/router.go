package Route

import (
	"face/Controller"
	"face/Middleware"
	"github.com/gin-gonic/gin"
	"github.com/mattn/go-colorable"
)

func InitRouter() {
	gin.ForceConsoleColor()
	//gin.New()
	//gin.SetMode(gin.ReleaseMode)
	//gin.SetMode(gin.DebugMode)
	gin.DefaultWriter = colorable.NewColorableStdout()
	router := gin.Default() //获取gin实例

	router.Use(Middleware.Cors()) //引入跨域中间件
	api := router.Group("api")    //路由组api
	{
		//用户登录相关
		api.POST("/Login", Controller.Login)
		api.GET("/AddUser", Controller.AddUser)
		api.POST("/Captcha", Controller.Captcha)

		api.POST("/GetIndexInfo", Controller.GetIndexInfo)
		api.GET("/GetUserInfo", Controller.GetUserInfo)

		api.POST("/SetClassInfo", Controller.SetClassInfo)
		api.POST("/GetClassList", Controller.GetClassList)
		api.POST("/DelClass", Controller.DelClass)

		api.POST("/GetUserList", Controller.GetUserList)
		api.POST("/SetUserInfo", Controller.SetUserInfo)

		api.POST("/GetLessonTime", Controller.GetLessonTime)
		api.POST("/SetLessonTime", Controller.SetLessonTime)
		api.POST("/DelLessonTime", Controller.DelLessonTime)

		api.POST("/GetLessonList", Controller.GetLessonList)

		api.POST("/GetClassRoomList", Controller.GetClassRoomList)
		api.POST("/SetClassRoom", Controller.SetClassRoom)
		api.POST("/DelClassRoom", Controller.DelClassRoom)
		api.POST("/Upload", Controller.Upload)

	}
	router.Run(":8887")
}
