package Models

type Face struct {
	Id   int    `gorm:"primary_key;AUTO_INCREMENT;type:int"`
	Name string `gorm:"column:name;type:varchar(255)"`
	Data []byte `gorm:"column:data;type:blob"`
}
