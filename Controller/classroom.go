package Controller

import (
	"face/Global"
	"face/Models"
	"face/Utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetClassRoomList(c *gin.Context) {
	type List struct {
		Name     string `json:"Name"`
		ID       string `json:"ID"`
		SignTime int    `json:"SignTime"`
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
	var ClassRoom []Models.Classroom

	var AllList []List

	tempsql := ""
	db := Global.DB
	if tempsql != "" {
		db.Where(tempsql).Offset((page - 1) * pageSize).Limit(pageSize).Find(&ClassRoom)
		db.Model(&Global.ClassRoomModel).Where(tempsql).Count(&Total)
	} else {
		db.Model(&Global.ClassRoomModel).Offset((page - 1) * pageSize).Limit(pageSize).Find(&ClassRoom)
		db.Model(&Global.ClassRoomModel).Count(&Total)
	}

	for _, record := range ClassRoom {

		list := []List{{ID: strconv.Itoa(record.Id), Name: record.Name, SignTime: record.Signtime}}
		AllList = append(AllList, list...)
	}
	c.JSON(http.StatusOK, LessonResp{
		0,
		Data{List: AllList, Total: int(Total), Page: page, PageSize: pageSize},
	})
}

func SetClassRoom(c *gin.Context) {
	data, _ := c.GetRawData()
	var ReqMap map[string]interface{}
	err := Utils.UnmarshalJSON(c, data, &ReqMap)
	if err != nil {
		return
	}

	//Id := ReqMap["ID"]
	Name := fmt.Sprintf("%v", ReqMap["Name"])
	Id := fmt.Sprintf("%v", ReqMap["ID"])
	Signtime, err := strconv.Atoi(fmt.Sprintf("%v", ReqMap["SignTime"]))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 7,
			"msg":  "时间格式错误",
		})
		return
	}

	var Msg string
	db := Global.DB

	if Id == "" {
		NewLessonTime := Models.Classroom{Name: Name, Signtime: Signtime}
		db.Create(&NewLessonTime)
		Msg = "新增成功"
	} else {
		db.Model(&Global.ClassRoomModel).Where("id = ?", Id).Updates(map[string]interface{}{"signtime": Signtime})
		Msg = "修改成功"
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  Msg,
	})

}

func DelClassRoom(c *gin.Context) {
	data, _ := c.GetRawData()
	var ReqMap map[string]string
	err := Utils.UnmarshalJSON(c, data, &ReqMap)
	if err != nil {
		return
	}
	ID := ReqMap["ID"]
	db := Global.DB
	db.Where("id = ?", ID).Delete(&Models.Classroom{})
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "删除成功",
	})
}
