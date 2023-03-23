package Controller

import (
	"encoding/base64"
	"face/Global"
	"face/Models"
	"face/Utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func HeartBeat(c *gin.Context) {
	//客户端心跳接口
	data, _ := c.GetRawData()
	var ReqMap map[string]int
	err := Utils.UnmarshalJSON(c, data, &ReqMap)
	if err != nil {
		return
	}

	Id := ReqMap["ID"]
	//fmt.Println(Id)
	//id代表教室id
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
	//c.JSON(http.StatusOK, gin.H{
	//	"code": 0,
	//	"msg":  ClassRoom.Name,
	//})

	//已知教室 则需要知道当前时间
	var Lesson []Models.Lesson
	NowTime := time.Now()
	loc, err := time.LoadLocation("Local")
	NowDate := NowTime.Format("2006-01-02")
	//NowTimes := NowTime.Format(" 15:04:05")
	//找到属于这个教室的课程
	//ClassRoom.TqSigntime
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
				"signtime":       BeforeTime.Unix() - NowTime.Unix(),
				"lessonname":     record.Name,
				"lessontimename": LessonTime.Name,
				"classroomname":  ClassRoom.Name,
			})
			return

		}
		//fmt.Println(LastTime)
		//fmt.Println(AfterTime)
		//fmt.Println(LessonT.Format("2006-01-02 15:04:05"))
		//fmt.Println(LessonT.Add(time.Duration(ClassRoom.TqSigntime)))

		//fmt.Println(LessonT.After())
		//fmt.Println(LessonT)
		//fmt.Println(LessonTime.Starttime)
	}
	c.JSON(http.StatusOK, gin.H{
		"code":          0,
		"classroomname": ClassRoom.Name,
		"cansign":       0,
	})
	//判断lessontime是否处于当前时间

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
	//User := "王小明"

	b, err := base64.StdEncoding.DecodeString(Image)
	nayoungFace, err := Global.FaceRe.RecognizeSingle(b)
	catID := Global.FaceRe.Classify(nayoungFace.Descriptor)
	var UserName string
	if catID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 2,
		})
		return
	} else {
		var User Models.UserList
		db.Model(&Global.UserListModel).Where("id=?", catID).Take(&User)
		UserName = User.UserName
	}
	var Total int64
	db.Model(&Global.SignHistoryModel).Where("signid=? and user=?", SignId, UserName).Count(&Total)
	if Total > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
		})
		return
	}
	NewSign := Models.Signhistory{Signid: SignId, User: UserName, Time: int(time.Now().Unix())}
	db.Create(&NewSign)
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"name": "王小明",
		"time": time.Now().Unix(),
	})

}
