package lobbysvc

import (
	"context"

	createlobby "github.com/philoj/goplanes/server/internal/app/handler/lobby/create"
	"github.com/philoj/goplanes/server/internal/domain/model"
)

var _ createlobby.LobbyService = &Service{}

type Service struct {
	repo LobbyRepository
}

func NewService(repo LobbyRepository) *Service {
	return &Service{
		repo: repo,
	}
}

type LobbyRepository interface {
	CreateLobby(ctx context.Context) (model.Lobby, error)
}

func (s *Service) CreateLobby(ctx context.Context) (model.Lobby, error) {
	return s.repo.CreateLobby(ctx)
}
