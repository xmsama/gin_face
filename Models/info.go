package Models

type Info struct {
	Id      int    `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	Date    string `gorm:"column:date;type:int(11)" json:"date"`
	Fail    int    `gorm:"column:fail;type:int(11)" json:"fail"`
	Success int    `gorm:"column:success;type:int(11)" json:"success"`
}

func (m *Info) TableName() string {
	return "info"
}
