package models

type User struct {
	ID        uint   `gorm:"primary" json:"id,omitempty"`
	IsAdmin   bool   `json:"is_admin,omitempty"`
	Login     string `json:"login,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Password  string `json:"password,omitempty" gorm:"->:false;<-:create"`
	Points    uint   `json:"points,omitempty"`
	Team      Team   `gorm:"embedded;embeddedPrefix:team_"`
	AvatarURI string `json:"avatar_uri,omitempty"`
}
