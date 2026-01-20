package lobbyrepo

import (
	"context"

	"github.com/philoj/goplanes/server/internal/domain/model"
	lobbysvc "github.com/philoj/goplanes/server/internal/domain/service/lobby"
)

var _ lobbysvc.LobbyRepository = &Repository{}

type Repository struct{}

type DB interface {
	Set(ctx context.Context)
}

func (r *Repository) CreateLobby(ctx context.Context) (model.Lobby, error) {
	//TODO implement me
	panic("implement me")
}
