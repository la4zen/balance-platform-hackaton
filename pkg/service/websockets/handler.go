package websockets

import (
	"github.com/gorilla/websocket"
	"github.com/la4zen/balance-platform-hackaton/pkg/models"
)

type Client struct {
	WS   *websocket.Conn
	User *models.User
}

type Handler struct {
	Clients map[uint]*Client
}

func NewHandler() *Handler {
	return &Handler{
		Clients: make(map[uint]*Client),
	}
}

func (h *Handler) SendEvent(event interface{}) {
	for _, client := range h.Clients {
		if client.WS.WriteJSON(event) != nil {
			delete(h.Clients, client.User.ID)
		}
	}
}
