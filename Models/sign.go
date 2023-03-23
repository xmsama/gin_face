package Models

type Sign struct {
	Id          int `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	Lessonid    int `gorm:"column:lessonid;type:int(11)" json:"lessonid"`
	Time        int `gorm:"column:time;type:int(11)" json:"time"`
	Classroomid int `gorm:"column:classroomid;type:int(11)" json:"classroomid"`
	Classid     int `gorm:"column:classid;type:int(11)" json:"classid"`
}

func (m *Sign) TableName() string {
	return "sign"
}
