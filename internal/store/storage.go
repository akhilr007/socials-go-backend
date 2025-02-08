package store

import (
	"context"
	"database/sql"
	"github.com/akhilr007/socials/internal/models"
)

type Storage struct {
	Posts interface {
		Create(ctx context.Context, post *models.Post) error
	}
	Users interface {
		Create(ctx context.Context) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts: &PostStore{db},
		Users: &UserStore{db},
	}
}
