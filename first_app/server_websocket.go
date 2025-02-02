package poker

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type playerServerWebSocket struct {
	*websocket.Conn
}

func newPlayerServerWebSocket(w http.ResponseWriter, r *http.Request) *playerServerWebSocket {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("problem upgrading connection to WebSocket %v\n", err)
	}

	return &playerServerWebSocket{conn}
}

func (w *playerServerWebSocket) WaitForMsg() string {
	_, msg, err := w.ReadMessage()
	if err != nil {
		log.Printf("error reading from WebSocket %v\n", err)
	}

	return string(msg)
}
