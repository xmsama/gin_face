package Global

import (
	"face/Models"
	"gorm.io/gorm"
)

var (
	DB           *gorm.DB
	FaceModel    Models.Face
	CaptchaModel Models.Captcha
	AccountModel Models.Account
	ClassModel   Models.Class
	JWTKey       string
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
