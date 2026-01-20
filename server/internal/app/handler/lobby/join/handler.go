package joinlobby

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var _ http.Handler = &Handler{}

type Handler struct {
	u *websocket.Upgrader
}

func NewHandler(upgrader *websocket.Upgrader) *Handler {
	return &Handler{u: upgrader}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	lobbyIDStr := r.PathValue("lobbyID")
	lobbyID, err := uuid.Parse(lobbyIDStr)
	if err != nil {
		slog.ErrorContext(r.Context(), "Invalid lobby ID", "lobbyID", lobbyIDStr, "err", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	playerIDStr := r.Header.Get("X-Player-ID")
	playerID, err := uuid.Parse(playerIDStr)
	if err != nil {
		slog.ErrorContext(r.Context(), "Invalid player ID", "playerID", playerIDStr, "err", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	conn, err := h.u.Upgrade(w, r, nil)
	if err != nil {
		slog.ErrorContext(r.Context(), "Failed to upgrade to websocket", "err", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Player %s joined lobby %s", playerID, lobbyID)))
	conn.Close()
}
