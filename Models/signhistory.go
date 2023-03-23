package Models

type Signhistory struct {
	Id     int    `gorm:"column:id;type:int(11);primary_key" json:"id"`
	Signid int    `gorm:"column:signid;type:int(11)" json:"signid"`
	User   string `gorm:"column:user;type:varchar(255)" json:"user"`
	Time   int    `gorm:"column:time;type:int(11)" json:"time"`
}

func (m *Signhistory) TableName() string {
	return "signhistory"
}
