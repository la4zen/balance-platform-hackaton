package store

import (
	"database/sql"
	"errors"

	"github.com/la4zen/balance-platform-hackaton/pkg/models"
)

func (s *Store) GetUser(u *models.User) error {
	r := s.DB.First(&u, "login=? and password=? OR id=?", u.Login, u.Password, u.ID)
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

func (s *Store) UpdateUser(u *models.User) error {
	r := s.DB.Save(&u)
	return r.Error
}

func (s *Store) GetLeaderBoardPos(u *models.User) (uint, error) {
	var result uint
	r := s.DB.Raw("SELECT index FROM (SELECT row_number() OVER () AS index, id as user_id FROM users ORDER BY points) AS rows WHERE rows.user_id = 0;").Scan(&result)
	if r.Error != nil {
		return 0, r.Error
	}
	return result, nil
}
