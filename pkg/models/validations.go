package models

import val "github.com/go-ozzo/ozzo-validation"

func (u *User) Validate() error {
	return val.ValidateStruct(u,
		val.Field(u.Login, val.Required, val.Length(4, 12)),
		val.Field(u.Password, val.Required),
		val.Field(u.FirstName, val.Required, val.Length(4, 12)),
		val.Field(u.LastName, val.Required, val.Length(4, 12)),
	)
}

func (t *Task) Validate() error {
	return val.ValidateStruct(
		val.Field(t.Title, val.Required),
		val.Field(t.Discription, val.Required),
	)
}
