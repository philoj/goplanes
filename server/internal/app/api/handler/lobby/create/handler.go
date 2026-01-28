package createlobby

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
	"github.com/philoj/goplanes/server/internal/domain/model"
)

var _ http.Handler = &Handler{}

type Handler struct {
	auth AuthService
	svc  LobbyService
}

func NewHandler(auth AuthService, svc LobbyService) *Handler {
	return &Handler{
		auth: auth,
		svc:  svc,
	}
}

type LobbyService interface {
	CreateLobby(ctx context.Context, ownerID uuid.UUID, name string) (model.Lobby, error)
}

type AuthService interface {
	GetUser(ctx context.Context) (model.User, error)
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user, err := h.auth.GetUser(r.Context())
	if err != nil {
		slog.ErrorContext(r.Context(), "Failed to get user", "err", err)
		w.WriteHeader(http.StatusForbidden)
		return
	}

	var req LobbyRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		slog.ErrorContext(r.Context(), "Failed to decode request body", "err", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	lobby, err := h.svc.CreateLobby(r.Context(), user.ID, req.Name)
	if err != nil {
		slog.ErrorContext(r.Context(), "Failed to create lobby", "err", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp := h.ImportLobby(lobby)
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		slog.ErrorContext(r.Context(), "Failed to write lobby", "err", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
