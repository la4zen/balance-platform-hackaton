package models

import "time"

type User struct {
	ID        uint `gorm:"primary" json:"id,omitempty"`
	IsAdmin   bool
	Login     string `json:"login,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Password  string `json:"password,omitempty"`
	Points    uint   `json:"points,omitempty"`
	AvatarURI string `json:"avatar_uri,omitempty"`
}

type Task struct {
	ID          uint `gorm:"primary"`
	Name        string
	Discription string
}

type TaskProgress struct {
	ID       uint `gorm:"primary"`
	User     User `gorm:"embedded;embeddedPrefix:user_"`
	Task     Task `gorm:"embedded;embeddedPrefix:task_"`
	Progress uint8
}
type Calls struct {
	ID        uint `gorm:"primary"`
	TimeStamp time.Time
	User      User `gorm:"embedded;embeddedPrefix:user_"`
}
