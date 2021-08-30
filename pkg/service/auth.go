package service

import (
	"github.com/la4zen/balance-platform-hackaton/pkg/models"
	"github.com/la4zen/balance-platform-hackaton/pkg/utils"
	"github.com/labstack/echo"
)

func (s *Service) Register(c echo.Context) error {
	user := &models.User{}
	c.Bind(user)
	if err := user.Validate(); err != nil {
		return err
	}
	if err := s.Store.CreateUser(user); err != nil {
		return err
	}
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return echo.ErrInternalServerError
	}
	return c.JSON(200, echo.Map{
		"token": token,
	})
}

func (s *Service) Login(c echo.Context) error {
	user := &models.User{}
	c.Bind(user)
	if err := user.Validate(); err != nil {
		return err
	}
	if err := s.Store.GetUser(user); err != nil {
		return err
	}
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return echo.ErrInternalServerError
	}
	return c.JSON(200, echo.Map{
		"token": token,
	})
}
