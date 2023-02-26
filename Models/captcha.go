package Models

type Captcha struct {
	Id     int    `gorm:"primary_key;AUTO_INCREMENT;type:int"`
	Base64 string `gorm:"column:base64;type:text"`
	Result string `gorm:"column:result;type:blob"`
	Time   int    `gorm:"column:time;"`
}
