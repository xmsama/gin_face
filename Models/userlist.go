package Models

type UserList struct {
	Id        int    `gorm:"primary_key;AUTO_INCREMENT"`
	UserName  string `gorm:"column:username;type:varchar(255)"`
	UserClass int    `gorm:"column:userclass;type:int(11)"`
	Image     string `gorm:"column:image;type:varchar(255)"`
	Face      []byte `gorm:"column:face;type:blob"`
	Time      int    `gorm:"column:time;type:int(11)"`
}

func (m *UserList) TableName() string {
	return "userlist"
}
