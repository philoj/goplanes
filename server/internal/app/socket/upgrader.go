package socket

import (
	"log/slog"
	"net/http"

	"github.com/gorilla/websocket"
)

func NewUpgrader() *websocket.Upgrader {
	return &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			slog.InfoContext(r.Context(), "Join request from origin", "origin", r.Header.Get("origin"))
			return r.Header.Get("origin") == "http://localhost:8081" // FIXME better origin value
		},
	}
}
