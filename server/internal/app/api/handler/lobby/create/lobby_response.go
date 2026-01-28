package createlobby

import "github.com/google/uuid"

type LobbyResponse struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	OwnerID uuid.UUID `json:"owner_id"`
}
