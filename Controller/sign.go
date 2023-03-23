package Controller

import (
	"face/Global"
	"face/Models"
	"face/Utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func GetSignHistory(c *gin.Context) {
	type List struct {
		LessonId    int    `json:"LessonId"`
		ID          int    `json:"ID"`
		ClassRoomId int    `json:"ClassRoomId"`
		ClassId     int    `json:"ClassId"`
		Time        string `json:"Time"`
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

	db := Global.DB
	var Sign []Models.Sign
	var Total int64
	tempsql := ""
	if tempsql != "" {
		db.Where(tempsql).Offset((page - 1) * pageSize).Limit(pageSize).Find(&Sign)
		db.Model(&Global.SignModel).Where(tempsql).Count(&Total)
	} else {
		db.Model(&Global.SignModel).Offset((page - 1) * pageSize).Limit(pageSize).Find(&Sign)
		db.Model(&Global.SignModel).Count(&Total)
	}
	var AllList []List

	for _, record := range Sign {
		t := time.Unix(int64(record.Time), 0)
		list := []List{{ID: record.Id, LessonId: record.Lessonid, ClassRoomId: record.Classroomid, ClassId: record.Classid, Time: t.Format("2006-01-02 15:04:05")}}
		AllList = append(AllList, list...)
	}
	c.JSON(http.StatusOK, LessonResp{
		0,
		Data{List: AllList, Total: int(Total), Page: page, PageSize: pageSize},
	})
	//Name= 课程名称
	//Id := ReqMap["ID"]
	//Name := fmt.Sprintf("%v", ReqMap["Name"])

}

func DelSignHistory(c *gin.Context) {
	data, _ := c.GetRawData()
	var ReqMap map[string]int
	err := Utils.UnmarshalJSON(c, data, &ReqMap)
	if err != nil {
		return
	}
	ID := ReqMap["ID"]
	db := Global.DB
	db.Where("id = ?", ID).Delete(&Models.Sign{})
	db.Where("signid = ?", ID).Delete(&Models.Signhistory{})
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "删除成功",
	})
}

func GetSignInfo(c *gin.Context) {
	type List struct {
		Name string `json:"Name"`
		Time string `json:"Time"`
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
	ID := fmt.Sprintf("%v", ReqMap["Id"])

	db := Global.DB
	var Sign []Models.Signhistory
	var Total int64
	tempsql := ""
	if tempsql != "" {
		db.Where(tempsql).Where("signid = ?", ID).Offset((page - 1) * pageSize).Limit(pageSize).Find(&Sign)
		db.Model(&Global.SignHistoryModel).Where(tempsql).Where("signid = ?", ID).Count(&Total)
	} else {
		db.Model(&Global.SignHistoryModel).Where("signid = ? ", ID).Offset((page - 1) * pageSize).Limit(pageSize).Find(&Sign)
		db.Model(&Global.SignHistoryModel).Where("signid = ?", ID).Count(&Total)
	}
	var AllList []List

	for _, record := range Sign {
		t := time.Unix(int64(record.Time), 0)
		list := []List{{Name: record.User, Time: t.Format("2006-01-02 15:04:05")}}
		AllList = append(AllList, list...)
	}
	c.JSON(http.StatusOK, LessonResp{
		0,
		Data{List: AllList, Total: int(Total), Page: page, PageSize: pageSize},
	})
}
