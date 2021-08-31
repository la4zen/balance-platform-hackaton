package service

import (
	"github.com/la4zen/balance-platform-hackaton/pkg/models"
	"github.com/labstack/echo"
)

func (s *Service) UpdatePhoto(c echo.Context) error {
	user := c.Get("user").(*models.User)
	c.Bind(user)
	if s.Store.UpdateUser(user) != nil {
		return echo.ErrInternalServerError
	}
	return c.NoContent(200)
}
