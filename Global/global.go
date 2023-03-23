package Global

import (
	"face/Models"
	"github.com/Kagami/go-face"
	"gorm.io/gorm"
)

var (
	DB               *gorm.DB
	FaceModel        Models.Face
	CaptchaModel     Models.Captcha
	AccountModel     Models.Account
	ClassModel       Models.Class
	UserListModel    Models.UserList
	LessontimeModel  Models.Lessontime
	LessonModel      Models.Lesson
	ClassRoomModel   Models.Classroom
	SignModel        Models.Sign
	SignHistoryModel Models.Signhistory
	JWTKey           string
	ImgPath          string
	FaceRe           *face.Recognizer
	//SettingModel      Models.Setting
	//CardModel         Models.Card
	//AccountModel      Models.Account
	//ProcessModel      Models.Process
	//Version           string
	//SexList           map[int]string
	//QueueList         map[int]string
	//ServerList        map[string]string
	//ServerListEn      map[string]string
	//StatusList        map[int]string
	//NowDownload       float64
	//NowDownloadName   string
	//NowDownloadMsg    string
	//NowDownloadStatus int
	//LastQueueTime     int64
	//QueueCount        int64
)
