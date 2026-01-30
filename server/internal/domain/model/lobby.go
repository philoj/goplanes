package model

import "github.com/google/uuid"

type Lobby struct {
	ID      uuid.UUID
	Name    string
	OwnerID uuid.UUID
}
