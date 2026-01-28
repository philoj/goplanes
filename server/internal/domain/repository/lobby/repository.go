package lobbyrepo

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/philoj/goplanes/server/internal/domain/model"
	lobbysvc "github.com/philoj/goplanes/server/internal/domain/service/lobby"
	"gorm.io/gorm"
)

var _ lobbysvc.LobbyRepository = &Repository{}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateLobby(ctx context.Context, ownerID uuid.UUID, name string) (model.Lobby, error) {
	lobby := Lobby{
		ID:      uuid.New(),
		Name:    name,
		OwnerID: ownerID,
	}
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		tErr := tx.Create(&lobby).Error
		if tErr != nil {
			return fmt.Errorf("inserting lobby: %w", tErr)
		}
		tErr = tx.Create(&LobbyMember{
			LobbyID: lobby.ID,
			UserID:  ownerID,
		}).Error
		if tErr != nil {
			return fmt.Errorf("inserting lobby member: %w", tErr)
		}
		return nil
	})
	if err != nil {
		return model.Lobby{}, err
	}
	return r.ExportLobby(lobby), nil
}
