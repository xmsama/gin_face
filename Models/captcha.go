package Models

type Captcha struct {
	Id     int    `gorm:"primary_key;AUTO_INCREMENT"`
	Base64 string `gorm:"column:base64;type:text"`
	Result string `gorm:"column:result;type:varchar(255)"`
	Time   int    `gorm:"column:time;"`
}
