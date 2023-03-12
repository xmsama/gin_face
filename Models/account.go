package Models

type Account struct {
	Id       int    `gorm:"primary_key;AUTO_INCREMENT"`
	Account  string `gorm:"column:account;type:varchar(255)"`
	Password string `gorm:"column:password;type:varchar(255)"`
	Name     string `gorm:"column:name;"`
}
