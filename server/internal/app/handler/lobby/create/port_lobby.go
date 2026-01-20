package createlobby

import "github.com/philoj/goplanes/server/internal/domain/model"

func (h *Handler) ImportLobby(l model.Lobby) LobbyResponse {
	return LobbyResponse{
		ID: l.ID,
	}
}
