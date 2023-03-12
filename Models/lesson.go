package Models

type Lesson struct {
	Id           int    `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT"`
	Name         string `gorm:"column:name;type:varchar(255)"`
	Classid      int    `gorm:"column:classid;type:int(11)"`
	Classroomid  int    `gorm:"column:classroomid;type:int(11)" `
	Lessontimeid int    `gorm:"column:lessontimeid;type:int(11)"`
}
