package Controller

import (
	"face/Global"
	"face/Models"
	"face/Utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"strconv"
)

func GetLessonTime(c *gin.Context) {
	type List struct {
		Name  string `json:"Name"`
		Stime string `json:"Stime"`
		ID    int    `json:"ID"`
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
		list := []List{{Name: record.Name, Stime: record.Starttime, ID: record.Id}}
		AllList = append(AllList, list...)
	}
	c.JSON(http.StatusOK, LessonTimeResp{
		0,
		Data{List: AllList, Total: int(Total), Page: page, PageSize: pageSize},
	})

}

func SetLessonTime(c *gin.Context) {
	data, _ := c.GetRawData()
	var ReqMap map[string]interface{}
	err := Utils.UnmarshalJSON(c, data, &ReqMap)
	if err != nil {
		return
	}
	//var Total int64
	var Msg string

	Stime := fmt.Sprintf("%v", ReqMap["Stime"])
	Name := fmt.Sprintf("%v", ReqMap["Name"])
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
	if Id == "" || Id == nil {
		NewLessonTime := Models.Lessontime{Name: Name, Starttime: Stime}
		db.Create(&NewLessonTime)
		Msg = "新增成功"
	} else {
		db.Model(&Global.LessontimeModel).Where("id = ?", Id).Updates(map[string]interface{}{"name": Name, "starttime": Stime})
		Msg = "修改成功"
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  Msg,
	})

}
func DelLessonTime(c *gin.Context) {
	data, _ := c.GetRawData()
	var ReqMap map[string]int
	err := Utils.UnmarshalJSON(c, data, &ReqMap)
	if err != nil {
		return
	}
	ID := ReqMap["ID"]
	db := Global.DB
	var Count int64
	db.Model(&Global.LessonModel).Where("lessontimeid = ? ", ID).Count(&Count)
	if Count > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 7,
			"msg":  "有课程绑定当前时间 请删除课程后再试！",
		})
		return
	}
	db.Where("id = ?", ID).Delete(&Models.Lessontime{})
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "删除成功",
	})
}

func GetLessonList(c *gin.Context) {
	type List struct {
		Name        string `json:"Name"`
		ClassId     int    `json:"ClassId"`
		ID          int    `json:"ID"`
		TimeId      int    `json:"TimeId"`
		ClassRoomId int    `json:"ClassRoomId"`
	}
	type Data struct {
		List     []List `json:"list"`
		Total    int    `json:"total"`
		PageSize int    `json:"pageSize"`
		Page     int    `json:"page"`
	}
	type LessonResp struct {
		Code int  `json:"code"`
		Data Data `json:"data"`
	}
	data, _ := c.GetRawData()
	var ReqMap map[string]interface{}
	err := Utils.UnmarshalJSON(c, data, &ReqMap)
	if err != nil {
		return
	}
	page, _ := strconv.Atoi(fmt.Sprintf("%v", ReqMap["page"]))
	pageSize, _ := strconv.Atoi(fmt.Sprintf("%v", ReqMap["pageSize"]))

	var Total int64
	var Lesson []Models.Lesson

	var AllList []List
	tempsql := Utils.SearchSql(ReqMap, 5)
	//tempsql := ""
	db := Global.DB
	if tempsql != "" {
		db.Where(tempsql).Offset((page - 1) * pageSize).Limit(pageSize).Find(&Lesson)
		db.Model(&Global.LessonModel).Where(tempsql).Count(&Total)
	} else {
		db.Model(&Global.LessonModel).Offset((page - 1) * pageSize).Limit(pageSize).Find(&Lesson)
		db.Model(&Global.LessonModel).Count(&Total)
	}

	for _, record := range Lesson {
		list := []List{{ID: record.Id, Name: record.Name, ClassId: record.Classid, ClassRoomId: record.Classroomid, TimeId: record.Lessontimeid}}
		AllList = append(AllList, list...)
	}
	c.JSON(http.StatusOK, LessonResp{
		0,
		Data{List: AllList, Total: int(Total), Page: page, PageSize: pageSize},
	})
}

func SetLessonInfo(c *gin.Context) {
	//{Name: "123123", ClassId: 5, ClassRoomId: "2", TimeId: "2"}
	data, _ := c.GetRawData()
	var Msg string
	var ReqMap map[string]interface{}
	err := Utils.UnmarshalJSON(c, data, &ReqMap)
	if err != nil {
		return
	}

	ClassId, _ := strconv.Atoi(fmt.Sprintf("%v", ReqMap["ClassId"]))
	ClassRoomId, _ := strconv.Atoi(fmt.Sprintf("%v", ReqMap["ClassRoomId"]))
	TimeId, _ := strconv.Atoi(fmt.Sprintf("%v", ReqMap["TimeId"]))
	Name := fmt.Sprintf("%v", ReqMap["Name"])

	Id := ReqMap["ID"]
	if ClassId == 0 || ClassRoomId == 0 || TimeId == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 7,
			"msg":  "请填写完整信息！",
		})
		return
	}
	db := Global.DB
	if Id == "" || Id == nil {
		//新增
		NewLesson := Models.Lesson{Name: Name, Classid: ClassId, Classroomid: ClassRoomId, Lessontimeid: TimeId}
		db.Create(&NewLesson)
		Msg = "新建成功"

	} else {
		//修改
		db.Model(&Global.LessonModel).Where("id = ?", Id).Updates(map[string]interface{}{"name": Name, "classid": ClassId, "classroomid": ClassRoomId,
			"lessontimeid": TimeId})
		Msg = "修改成功"
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  Msg,
	})

}
func DelLesson(c *gin.Context) {
	data, _ := c.GetRawData()
	var ReqMap map[string]int
	err := Utils.UnmarshalJSON(c, data, &ReqMap)
	if err != nil {
		return
	}
	ID := ReqMap["ID"]
	db := Global.DB
	db.Where("id = ?", ID).Delete(&Models.Lesson{})
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "删除成功",
	})
}
