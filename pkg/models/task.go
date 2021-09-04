package models

type Task struct {
	ID          uint   `gorm:"primary" json:"id,omitempty"`
	Title       string `json:"name,omitempty"`
	Discription string `json:"discription,omitempty"`
}

type TaskProgress struct {
	ID       uint  `gorm:"primary" json:"id,omitempty"`
	User     User  `gorm:"embedded;embeddedPrefix:user_" json:"user,omitempty"`
	Task     Task  `gorm:"embedded;embeddedPrefix:task_" json:"task,omitempty"`
	Progress uint8 `json:"progress,omitempty"`
}
