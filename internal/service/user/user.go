package user

import (
	"context"
	"github.com/SavinDevelop/techcrm-go/internal/model"
	"github.com/SavinDevelop/techcrm-go/internal/repository/user"
	"github.com/google/uuid"
	"time"
)

type Service struct {
	repo *user.Repository
}

func NewService(repo *user.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, user model.User) error {
	user.ID = uuid.New().String()
	user.CreatedAt = time.Now().Format(time.RFC3339)
	user.UpdateAt = user.CreatedAt
	return s.repo.Create(ctx, user)
}
