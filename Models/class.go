package Models

type Class struct {
	Id   int    `gorm:"primary_key;AUTO_INCREMENT"`
	Name string `gorm:"column:name;type:varchar(255)"`
	Time int    `gorm:"column:time;"`
}
