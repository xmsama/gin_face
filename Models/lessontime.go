package Models

type Lessontime struct {
	Id        int    `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT"`
	Name      string `gorm:"column:name;type:varchar(255)"`
	Starttime string `gorm:"column:starttime;type:varchar(255)" `
}

func (m *Lessontime) TableName() string {
	return "lessontime"
}
