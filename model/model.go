package model

type User struct {
	ID        uint   `gorm:"primary_key,auto_increment"`
	Phone     string `gorm:"varchar(20),unique" binding:"phone"`
	Account   string `gorm:"varchar(50),unique" binding:"required,max=50" `
	Password  string `gorm:"varchar(30)" binding:"required,min=6,max=30"`
	Nick      string `gorm:"varchar(50)" binding:"max=50"`
	Address   string `gorm:"varchar(255)"`
	Age       uint8
	Introduce string `gorm:"text"`
}
