package models

type Check struct {
	ID    uint   `json:"id,omitempty" gorm:"primary"`
	INN   string `json:"inn,omitempty"`
	SNILS string `json:"snils,omitempty"`
}
