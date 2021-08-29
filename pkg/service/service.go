package service

import (
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

func (s *Service) Accessible(c echo.Context) error {
	return c.NoContent(200)
}
