package service

import (
	"github.com/la4zen/balance-platform-hackaton/pkg/store"
)

type Service struct {
	Store *store.Store
}

func NewService(s *store.Store) *Service {
	return &Service{
		Store: s,
	}
}
