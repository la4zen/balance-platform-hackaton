package models

type Team struct {
	ID   uint `gorm:"primary"`
	Name string
}
