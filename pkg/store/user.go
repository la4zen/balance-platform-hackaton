package store

import (
	"database/sql"
	"errors"

	"github.com/la4zen/balance-platform-hackaton/pkg/models"
)

func (s *Store) GetUser(u *models.User) error {
	r := s.DB.First(&u, "login=? and password=?", u.Login, u.Password)
	if r.Error != nil {
		if r.Error == sql.ErrNoRows {
			return errors.New("user not found")
		}
		return r.Error
	}
	return nil
}

func (s *Store) CreateUser(u *models.User) error {
	r := s.DB.First(&u, "login=?", u.Login)
	if r.RowsAffected > 0 {
		return errors.New("user with this login already exists")
	}
	r = s.DB.Create(&u)
	return r.Error
}
