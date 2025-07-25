package user

import (
	"context"
	"github.com/SavinDevelop/techcrm-go/internal/model"
	"github.com/SavinDevelop/techcrm-go/pkg/db"
)

type Repository struct {
	db *db.Postgres
}

func NewRepository(db *db.Postgres) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, user model.User) error {
	query := `INSERT INTO users (id, email, password, is_active, created_at, update_at) 
              VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.Exec(ctx, query, user.ID, user.Email, user.Password, user.IsActive, user.CreatedAt, user.UpdateAt)
	return err
}
