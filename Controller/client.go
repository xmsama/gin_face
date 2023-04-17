package Controller

import (
	"encoding/base64"
	"face/Global"
	"face/Models"
	"face/Utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func HeartBeat(c *gin.Context) {
	data, _ := c.GetRawData()
	var ReqMap map[string]int
	err := Utils.UnmarshalJSON(c, data, &ReqMap)
	if err != nil {
		return
	}
	Id := ReqMap["ID"]
	var ClassRoom Models.Classroom
	db := Global.DB
	db.Model(&Global.ClassRoomModel).Where("id = ? ", Id).Take(&ClassRoom)
	if ClassRoom.Name == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 7,
			"msg":  "ID不存在！",
		})
		return
	}
	//已知教室 则需要知道当前时间
	var Lesson []Models.Lesson
	NowTime := time.Now()
	loc, err := time.LoadLocation("Local")
	NowDate := NowTime.Format("2006-01-02")
	db.Model(&Global.LessonModel).Where("classroomid = ?", Id).Find(&Lesson)
	for _, record := range Lesson {
		var LessonTime Models.Lessontime
		db.Model(&Global.LessontimeModel).Where("id = ? ", record.Lessontimeid).Take(&LessonTime)
		//拼接时间
		LessonT, _ := time.ParseInLocation("2006-01-02 15:04:05", NowDate+" "+LessonTime.Starttime, loc)
		AfterTime := LessonT.Add(-(time.Second * time.Duration(ClassRoom.TqSigntime)))
		BeforeTime := LessonT.Add(time.Second * time.Duration(ClassRoom.Signtime))
		if NowTime.After(AfterTime) && NowTime.Before(BeforeTime) {
			//创建签到表

			NewSign := Models.Sign{Lessonid: record.Lessontimeid, Classroomid: ClassRoom.Id, Classid: record.Classid, Time: int(time.Now().Unix())}
			db.Create(&NewSign)
			c.JSON(http.StatusOK, gin.H{
				"code":           0,
				"cansign":        1,
				"signid":         NewSign.Id,
				"signtime":       NowTime.Unix() + (BeforeTime.Unix() - NowTime.Unix()),
				"lessonname":     record.Name,
				"lessontimename": LessonTime.Name,
				"classroomname":  ClassRoom.Name,
			})
			return

		}

	}
	c.JSON(http.StatusOK, gin.H{
		"code":          0,
		"classroomname": ClassRoom.Name,
		"cansign":       0,
	})
}
func GetRandom() string {
	rand.Seed(time.Now().UnixNano()) // 初始化随机数生成器
	var builder strings.Builder
	for i := 0; i < 5; i++ {
		ran := "abcdefghjiklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"[rand.Intn(62)]
		builder.WriteByte(ran)
	}
	temp := builder.String()
	return temp
}
func Detected(c *gin.Context) {
	data, _ := c.GetRawData()
	var ReqMap map[string]interface{}
	err := Utils.UnmarshalJSON(c, data, &ReqMap)
	if err != nil {
		return
	}
	db := Global.DB
	SignId := int(ReqMap["SignId"].(float64))
	Image := ReqMap["Image"].(string)
	b, _ := base64.StdEncoding.DecodeString(Image)
	if err != nil {
		fmt.Println("b64Error:", err)
	}
	FileName := GetRandom() + strconv.Itoa(int(time.Now().Unix())) + ".jpg"
	//fmt.Println(FileName)
	file, err := os.Create(Global.ImgPath + "/temp/" + FileName)
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
	DFace, err := Global.FaceRe.RecognizeFile(Global.ImgPath + "/temp/" + FileName)
	if err != nil {
		fmt.Println("RecognizeSingle:", err)
	}
	var Info Models.Info
	NowDay := time.Now().Format("2006-01-02")
	db.Where("date= ?", NowDay).Take(&Info)
	if Info.Date == "" {
		//不存在这个日期
		db.Create(&Models.Info{
			Date:    NowDay,
			Success: 0,
			Fail:    0,
		})
	}
	var SignSuccess string
	for _, record := range DFace {
		catID := Global.FaceRe.ClassifyThreshold(record.Descriptor, 0.4)
		if catID == 0 {
			db.Model(&Global.InfoModel).Where("date= ?", NowDay).Updates(map[string]interface{}{"fail": gorm.Expr("fail + 1")})
			continue
		}
		var User Models.UserList
		db.Model(&Global.UserListModel).Where("id=?", catID).Take(&User)
		var Total int64
		db.Model(&Global.SignHistoryModel).Where("signid=? and user=?", SignId, User.UserName).Count(&Total)
		if Total < 1 {
			NewSign := Models.Signhistory{Signid: SignId, User: User.UserName, Time: int(time.Now().Unix())}
			db.Create(&NewSign)
		}
		db.Model(&Global.InfoModel).Where("date= ?", NowDay).Updates(map[string]interface{}{"success": gorm.Expr("success + 1")})
		SignSuccess += User.UserName + "|"
	}
	os.Remove(Global.ImgPath + "/temp/" + FileName)
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"name": SignSuccess,
		"time": time.Now().Unix(),
	})

}
