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
	if tempsql != "" {
		db.Where(tempsql).Offset((page - 1) * pageSize).Limit(pageSize).Order("time desc").Find(&Class)
		db.Model(&Global.ClassModel).Where(tempsql).Count(&Total)
	} else {
		db.Model(&Global.ClassModel).Offset((page - 1) * pageSize).Limit(pageSize).Order("time desc").Find(&Class)
		db.Model(&Global.ClassModel).Count(&Total)
	}
	var AllList []List
	for _, record := range Class {
		var Count int64
		db.Model(&Global.UserListModel).Where("userclass=?", record.Id).Count(&Count)
		list := []List{{Name: record.Name, Count: int(Count), ID: record.Id, Status: "正常", Time: int(time.Now().Unix())}}
		AllList = append(AllList, list...)
	}
	c.JSON(http.StatusOK, ClassListResp{
		0,
		Data{List: AllList, Total: int(Total), Page: page, PageSize: pageSize},
	})
}

func Upload(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
	})
}
