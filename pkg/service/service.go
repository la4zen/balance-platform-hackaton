package service

import (
	"github.com/la4zen/balance-platform-hackaton/pkg/models"
	"github.com/la4zen/balance-platform-hackaton/pkg/store"
	"github.com/labstack/echo"
)

type Service struct {
	Store *store.Store
}

func NewService(s *store.Store) *Service {
	return &Service{
		Store: s,
	}
}

func (s *Service) Register(c echo.Context) error {
	user := &models.User{}
	c.Bind(user)
	if err := user.Validate(); err != nil {
		return err
	}
	if err := s.Store.CreateUser(user); err != nil {
		return err
	}
	// TODO : generate tokens
	return c.NoContent(200)
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
	// TODO : generate tokens
	return c.NoContent(200)
}
