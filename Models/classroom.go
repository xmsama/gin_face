package Models

type Classroom struct {
	Id       int    `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT"`
	Name     string `gorm:"column:name;type:varchar(255)"`
	Signtime int    `gorm:"column:signtime;type:int(11);default:300;"`
}

func (m *Classroom) TableName() string {
	return "classroom"
}
