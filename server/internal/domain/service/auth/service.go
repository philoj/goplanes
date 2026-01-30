package authsvc

import (
	"context"
	"net/http"

	"github.com/philoj/goplanes/server/internal/app/api"
	"github.com/philoj/goplanes/server/internal/domain/model"
)

var _ api.AuthService = &Service{}

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Middleware(next http.Handler) http.Handler {
	//TODO implement me
	panic("implement me")
}

func (s *Service) GetUser(ctx context.Context) (model.User, error) {
	//TODO implement me
	panic("implement me")
}
