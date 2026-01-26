package lobbyrepo

import (
	"context"

	"github.com/philoj/goplanes/server/internal/domain/model"
	lobbysvc "github.com/philoj/goplanes/server/internal/domain/service/lobby"
	"gorm.io/gorm"
)

var _ lobbysvc.LobbyRepository = &Repository{}

type Repository struct {
	db *gorm.DB
}

type DB interface {
	Set(ctx context.Context)
}

func (r *Repository) CreateLobby(ctx context.Context) (model.Lobby, error) {

}
