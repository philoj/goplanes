package createlobby

import "github.com/google/uuid"

type LobbyResponse struct {
	ID uuid.UUID `json:"id"`
}
