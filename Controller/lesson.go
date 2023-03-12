package Controller

import (
	"face/Global"
	"face/Models"
	"face/Utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"strconv"
)

func GetLessonTime(c *gin.Context) {
	type List struct {
		Name  string `json:"Name"`
		Stime string `json:"Stime"`
		ID    string `json:"ID"`
	}
	type Data struct {
		List     []List `json:"list"`
		Total    int    `json:"total"`
		PageSize int    `json:"pageSize"`
		Page     int    `json:"page"`
	}
	type LessonTimeResp struct {
		Code int  `json:"code"`
		Data Data `json:"data"`
	}

	db := Global.DB
	var LessonTime []Models.Lessontime
	var AllList []List
	db.Order("starttime").Find(&LessonTime)
	data, _ := c.GetRawData()
	var datamap map[string]int
	var Total int64
	err := Utils.UnmarshalJSON(c, data, &datamap)
	if err != nil {
		return
	}
	page, _ := datamap["page"]
	pageSize, _ := datamap["pageSize"]
	db.Model(&Global.LessontimeModel).Count(&Total)
	for _, record := range LessonTime {
		list := []List{{Name: record.Name, Stime: record.Starttime, ID: strconv.Itoa(record.Id)}}
		AllList = append(AllList, list...)
	}
	c.JSON(http.StatusOK, LessonTimeResp{
		0,
		Data{List: AllList, Total: int(Total), Page: page, PageSize: pageSize},
	})

}

func SetLessonTime(c *gin.Context) {
	data, _ := c.GetRawData()
	var ReqMap map[string]string
	err := Utils.UnmarshalJSON(c, data, &ReqMap)
	if err != nil {
		return
	}
	//var Total int64
	var Msg string
	Name := ReqMap["Name"]
	Stime := ReqMap["Stime"]
	Id := ReqMap["ID"]
	pattern := `^([01][0-9]|2[0-4]):([0-5][0-9]):([0-5][0-9])$`
	re := regexp.MustCompile(pattern)
	if !re.MatchString(Stime) {
		c.JSON(http.StatusOK, gin.H{
			"code": 7,
			"msg":  "时间格式错误",
		})
		return
	}

	db := Global.DB
	if Id == "" {
		NewLessonTime := Models.Lessontime{Name: Name, Starttime: Stime}
		db.Create(&NewLessonTime)
		Msg = "新增成功"
	} else {
		db.Model(&Global.LessontimeModel).Where("id = ?", Id).Updates(map[string]interface{}{"starttime": Stime})
		Msg = "修改成功"
	}
	//db.Model(&Global.LessontimeModel).Where("Name = ?", Name).Count(&Total)
	////if Total < 1 {
	////	NewLessonTime := Models.Lessontime{Name: Name, Starttime: Stime}
	////	db.Create(&NewLessonTime)
	////	Msg = "新增成功"
	//} else if Total > 0 {
	//	db.Model(&Global.LessontimeModel).Where("Id = ?", Id).Updates(map[string]interface{}{"Stime": Stime})
	//	Msg = "修改成功"
	//} else {
	//	c.JSON(http.StatusOK, gin.H{
	//		"code": 7,
	//		"msg":  "异常错误",
	//	})
	//	return
	//}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  Msg,
	})

}
func DelLessonTime(c *gin.Context) {
	data, _ := c.GetRawData()
	var ReqMap map[string]string
	err := Utils.UnmarshalJSON(c, data, &ReqMap)
	if err != nil {
		return
	}
	ID := ReqMap["ID"]
	db := Global.DB
	db.Where("id = ?", ID).Delete(&Models.Lessontime{})
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "删除成功",
	})
}

func GetLessonList(c *gin.Context) {

}
