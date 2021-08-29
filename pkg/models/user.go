package models

type User struct {
	ID       uint `gorm:"primary"`
	Login    string
	Password string
}
