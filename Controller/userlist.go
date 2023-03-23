package Controller

import (
	"encoding/base64"
	"face/Global"
	"face/Models"
	"face/Utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	"strings"
)

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
		var Status string
		if record.Face != nil {
			Status = "已生成人脸数据"
		} else {
			Status = "未生成人脸数据"
		}

		list := []List{{ID: record.Id, Name: record.UserName, Class: Class.Name, ClassId: record.UserClass, Image: Utils.ReadImage(record.Image), Status: Status}}
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
		//fmt.Println(strings.Replace(B64, "data:image/jpeg;base64,", "", 1))
		b64data := B64[strings.IndexByte(B64, ',')+1:]
		//DB64 := strings.Replace(B64, "data:image/jpeg;base64,", "", 1)
		b, err := base64.StdEncoding.DecodeString(b64data)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusOK, gin.H{
				"code": 7,
				"msg":  "图片格式错误！",
			})
			return
		}
		sizeInKB := len(b) / 1024
		if sizeInKB > 1024 {
			c.JSON(http.StatusOK, gin.H{
				"code": 7,
				"msg":  "图片超过1M！",
			})
			return
		}

		NewLessonTime := Models.UserList{UserName: Name, UserClass: ClassId}
		db.Create(&NewLessonTime)

		// 将解码后的数据写入文件
		file, err := os.Create(Global.ImgPath + "/" + strconv.Itoa(NewLessonTime.Id) + ".jpg")
		if err != nil {
			fmt.Println("创建文件失败", err)
			return
		}
		defer file.Close()
		_, err = file.Write(b)
		if err != nil {
			fmt.Println("写入文件失败")
			return
		}
		Msg = "新增成功"
		Utils.AddFace(NewLessonTime.Id, b)

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
func DelUser(c *gin.Context) {
	data, _ := c.GetRawData()
	var ReqMap map[string]int
	err := Utils.UnmarshalJSON(c, data, &ReqMap)
	if err != nil {
		return
	}
	ID := ReqMap["ID"]
	db := Global.DB
	db.Where("id = ?", ID).Delete(&Models.UserList{})
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "删除成功",
	})
}
