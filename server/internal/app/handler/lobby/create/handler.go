package createlobby

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/philoj/goplanes/server/internal/domain/model"
)

var _ http.Handler = &Handler{}

type Handler struct {
	svc LobbyService
}

func NewHandler(svc LobbyService) *Handler {
	return &Handler{
		svc: svc,
	}
}

type LobbyService interface {
	CreateLobby(ctx context.Context) (model.Lobby, error)
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	lobby, err := h.svc.CreateLobby(r.Context())
	if err != nil {
		slog.ErrorContext(r.Context(), "Failed to create lobby", "err", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(h.ImportLobby(lobby))
	if err != nil {
		slog.ErrorContext(r.Context(), "Failed to write lobby", "err", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
