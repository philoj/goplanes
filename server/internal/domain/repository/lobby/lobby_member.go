package lobbyrepo

import "github.com/google/uuid"

type LobbyMember struct {
	// TODO fill in tags
	LobbyID uuid.UUID
	UserID  uuid.UUID
}
