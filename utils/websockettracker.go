package utils

import (
	"github.com/gorilla/websocket"
)

type WebSocketManager interface {
	Register(ws *websocket.Conn)

	Write(msg []byte) (n int, e error)
}

type wsManager struct {
	sockets []websocket.Conn
}

func CreateWSManager() WebSocketManager {
	return &wsManager{sockets: make([]websocket.Conn, 0)}
}

func (ws *wsManager) Register(conn *websocket.Conn) {
	ws.sockets = append(ws.sockets, *conn)
}

func (ws *wsManager) Write(msg []byte) (n int, e error) {
	invalid := make([]int, 0)
	for k, v := range ws.sockets {
		err := v.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			invalid = append(invalid, k)
		}
	}
	if len(invalid) > 0 {
		for b := range invalid {
			ws.sockets = append(ws.sockets[:b], ws.sockets[b+1:]...)
		}
	}
	n = len(msg)
	return
}
