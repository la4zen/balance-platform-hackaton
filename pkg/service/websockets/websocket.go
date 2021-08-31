package websockets

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/la4zen/balance-platform-hackaton/pkg/models"
	"github.com/labstack/echo"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func (h *Handler) WSHandler(c echo.Context) error {
	client := Client{
		User: c.Get("user").(*models.User),
	}
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	client.WS = ws
	go func() {
		defer ws.Close()
		var data map[string]interface{}
		for {
			if ws.ReadJSON(&data) != nil {
				return
			}
			switch {
			}
		}
	}()
	return nil
}
