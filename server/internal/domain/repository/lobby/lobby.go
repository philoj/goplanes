package lobbyrepo

import (
	"github.com/google/uuid"
	"github.com/philoj/goplanes/server/internal/domain/model"
)

type Lobby struct {
	// TODO fill in tags
	ID      uuid.UUID
	Name    string
	OwnerID uuid.UUID
}

func (r *Repository) ExportLobby(l Lobby) model.Lobby {
	return model.Lobby{
		ID:      l.ID,
		Name:    l.Name,
		OwnerID: l.OwnerID,
	}
}
