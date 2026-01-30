package lobbysvc

import (
	"context"

	"github.com/google/uuid"
	"github.com/philoj/goplanes/server/internal/app/api/handler/lobby/create"
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
	CreateLobby(ctx context.Context, ownerID uuid.UUID, name string) (model.Lobby, error)
}

func (s *Service) CreateLobby(ctx context.Context, ownerID uuid.UUID, name string) (model.Lobby, error) {
	return s.repo.CreateLobby(ctx, ownerID, name)
}
