package handler

import (
	"log"
	"log/slog"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gorilla/websocket"
)

type LobbyHandler struct {
	u   *websocket.Upgrader
	svc LobbyService
}

func NewLobbyHandler(u *websocket.Upgrader, svc LobbyService) *LobbyHandler {
	return &LobbyHandler{u: u, svc: svc}
}

type LobbyService interface {
	CreatePlayer(id int, conn *websocket.Conn)
}

func (h *LobbyHandler) Handle(w http.ResponseWriter, r *http.Request) {
	log.Println("socket hit", r.URL.Path)
	id, err := strconv.Atoi(regexp.MustCompile("/lobby/(\\d+)").FindStringSubmatch(r.URL.Path)[1])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	conn, err := h.u.Upgrade(w, r, nil)
	if err != nil {
		slog.ErrorContext(r.Context(), "Failed to upgrade to websocket", "err", err)
		return
	}
	h.svc.CreatePlayer(id, conn)
}
