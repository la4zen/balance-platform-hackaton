package service

import (
	"github.com/la4zen/balance-platform-hackaton/pkg/service/websockets"
	"github.com/la4zen/balance-platform-hackaton/pkg/store"
)

type Service struct {
	Store     *store.Store
	WSHandler *websockets.Handler
}

func NewService(s *store.Store) *Service {
	return &Service{
		Store:     s,
		WSHandler: websockets.NewHandler(),
	}
}
