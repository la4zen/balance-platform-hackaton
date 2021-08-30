package models

type User struct {
	ID        uint   `gorm:"primary" json:"id,omitempty"`
	Login     string `json:"login,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Password  string `json:"password,omitempty"`
	Points    uint   `json:"points,omitempty"`
	AvatarURI string `json:"avatar_uri,omitempty"`
}
