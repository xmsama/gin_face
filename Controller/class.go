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

func SetClassInfo(c *gin.Context) {
	data, _ := c.GetRawData()
	var ReqMap map[string]string
	err := Utils.UnmarshalJSON(c, data, &ReqMap)
	if err != nil {
		return
	}
	Name := ReqMap["Name"]
	Id := ReqMap["ID"]
	db := Global.DB
	var Msg string

	if Id == "" {
		NewClass := Models.Class{Name: Name}
		db.Create(&NewClass)
		Msg = "新增成功"
	} else {
		db.Model(&Global.LessontimeModel).Where("id = ?", Id).Updates(map[string]interface{}{"name": Name})
		Msg = "修改成功"
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  Msg,
	})
}

func DelClass(c *gin.Context) {
	data, _ := c.GetRawData()
	var ReqMap map[string]int
	err := Utils.UnmarshalJSON(c, data, &ReqMap)
	if err != nil {
		return
	}
	ID := ReqMap["ID"]
	db := Global.DB
	db.Where("id = ?", ID).Delete(&Models.Class{})
	db.Where("userclass = ?", ID).Delete(&Models.UserList{})
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "删除成功",
	})
}

func GetClassList(c *gin.Context) {
	type List struct {
		Name   string `json:"Name"`
		Time   int    `json:"Time"`
		Count  int    `json:"Count"`
		Status string `json:"Status"`
		ID     int    `json:"ID"`
	}
	type Data struct {
		List     []List `json:"list"`
		Total    int    `json:"total"`
		PageSize int    `json:"pageSize"`
		Page     int    `json:"page"`
	}
	type ClassListResp struct {
		Code int  `json:"code"`
		Data Data `json:"data"`
	}
	db := Global.DB
	var Class []Models.Class
	data, _ := c.GetRawData()
	var datamap map[string]interface{}
	err := Utils.UnmarshalJSON(c, data, &datamap)
	if err != nil {
		return
	}
	tempsql := Utils.SearchSql(datamap, 5)
	page, _ := strconv.Atoi(fmt.Sprintf("%v", datamap["page"]))
	pageSize, _ := strconv.Atoi(fmt.Sprintf("%v", datamap["pageSize"]))
	var Total int64
	//fmt.Println(page)
	//fmt.Println(pageSize)
	//fmt.Println(Total)
	if tempsql != "" {
		db.Where(tempsql).Offset((page - 1) * pageSize).Limit(pageSize).Order("time desc").Find(&Class)
		db.Model(&Global.ClassModel).Where(tempsql).Count(&Total)
	} else {
		db.Model(&Global.ClassModel).Offset((page - 1) * pageSize).Limit(pageSize).Order("time desc").Find(&Class)
		db.Model(&Global.ClassModel).Count(&Total)
	}
	var AllList []List
	for _, record := range Class {
		list := []List{{Name: record.Name, Count: 114, ID: record.Id, Status: "正常", Time: int(time.Now().Unix())}}
		AllList = append(AllList, list...)
	}

	c.JSON(http.StatusOK, ClassListResp{
		0,
		Data{List: AllList, Total: int(Total), Page: page, PageSize: pageSize},
	})
	//c.JSON(http.StatusOK, gin.H{
	//	"code": 0,
	//	"msg":  "giao",
	//})
}
func GetUserList(c *gin.Context) {
	type List struct {
		Name    string `json:"Name"`
		ID      int    `json:"ID"`
		Class   string `json:"Class"`
		ClassId int    `json:"ClassId"`
		Image   string `json:"Image"`
		Status  string `json:"Status"`
	}
	type Data struct {
		List     []List `json:"list"`
		Total    int    `json:"total"`
		PageSize int    `json:"pageSize"`
		Page     int    `json:"page"`
	}
	type UserListResp struct {
		Code int  `json:"code"`
		Data Data `json:"data"`
	}

	db := Global.DB
	var User []Models.UserList
	data, _ := c.GetRawData()
	var datamap map[string]interface{}
	err := Utils.UnmarshalJSON(c, data, &datamap)
	if err != nil {
		return
	}
	tempsql := Utils.SearchSql(datamap, 5)
	page, _ := strconv.Atoi(fmt.Sprintf("%v", datamap["page"]))
	pageSize, _ := strconv.Atoi(fmt.Sprintf("%v", datamap["pageSize"]))
	var Total int64
	if tempsql != "" {
		db.Where(tempsql).Offset((page - 1) * pageSize).Limit(pageSize).Order("time desc").Find(&User)
		db.Model(&Global.UserListModel).Where(tempsql).Count(&Total)
	} else {
		db.Model(&Global.UserListModel).Offset((page - 1) * pageSize).Limit(pageSize).Order("time desc").Find(&User)
		db.Model(&Global.UserListModel).Count(&Total)
	}
	var AllList []List
	for _, record := range User {
		var Class Models.Class
		db.Where("id = ? ", record.UserClass).Take(&Class)
		list := []List{{ID: record.Id, Name: record.UserName, Class: Class.Name, ClassId: record.UserClass, Image: record.Image, Status: "未生成人脸数据"}}
		AllList = append(AllList, list...)
	}

	c.JSON(http.StatusOK, UserListResp{
		0,
		Data{List: AllList, Total: int(Total), Page: page, PageSize: pageSize},
	})
}

func SetUserInfo(c *gin.Context) {
	data, _ := c.GetRawData()
	var ReqMap map[string]interface{}
	err := Utils.UnmarshalJSON(c, data, &ReqMap)
	if err != nil {
		return
	}
	Name := fmt.Sprintf("%v", ReqMap["Name"])
	ClassId, _ := strconv.Atoi(fmt.Sprintf("%v", ReqMap["ClassId"]))
	Id := ReqMap["ID"]
	B64 := fmt.Sprintf("%v", ReqMap["Image"])
	if ReqMap["ClassId"] == "" || Name == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 7,
			"msg":  "信息不全！",
		})
		return
	}
	var Msg string
	db := Global.DB
	if Id == "" || Id == nil {
		if B64 == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 7,
				"msg":  "请上传图片！",
			})
			return
		}
		NewLessonTime := Models.UserList{UserName: Name, UserClass: ClassId, Image: B64}
		db.Create(&NewLessonTime)
		Msg = "新增成功"
	} else {
		if B64 == "" {
			db.Model(&Global.UserListModel).Where("id = ?", Id).Updates(map[string]interface{}{"username": Name, "userclass": ClassId})
		} else {
			db.Model(&Global.UserListModel).Where("id = ?", Id).Updates(map[string]interface{}{"username": Name, "userclass": ClassId, "image": B64})
		}

		Msg = "修改成功"
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  Msg,
	})

}

func Upload(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
	})
}
