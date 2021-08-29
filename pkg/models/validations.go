package models

import val "github.com/go-ozzo/ozzo-validation"

func (u *User) Validate() error {
	return val.ValidateStruct(u,
		val.Field(u.Login, val.Required),
		val.Field(u.Password, val.Required),
	)
}
