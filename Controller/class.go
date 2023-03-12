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
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "giao",
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
		list := []List{{Name: record.Name, Count: 114, Status: "正常", Time: int(time.Now().Unix())}}
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
		Name   string `json:"Name"`
		Class  string `json:"Class"`
		Image  string `json:"Image"`
		Status string `json:"Status"`
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
		list := []List{{Name: record.UserName, Class: "我是一个班级个班级个班级个班级个班级", Image: "/img/abc.png", Status: "未生成人脸数据"}}
		AllList = append(AllList, list...)
	}

	c.JSON(http.StatusOK, UserListResp{
		0,
		Data{List: AllList, Total: int(Total), Page: page, PageSize: pageSize},
	})
}
