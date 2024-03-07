package storage

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/kerucko/auth/internal/models"
)

var (
	ErrUserNotFound = fmt.Errorf("user not found")
	ErrUserExists   = fmt.Errorf("user already exists")
	ErrAppNotFound  = fmt.Errorf("app not found")
)

type Storage struct {
	db *sql.DB
}

func NewStorage(host string, port int, dbname string, user string, password string, timeout string) *Storage {
	return &Storage{}
}

func (s *Storage) SaveUser(ctx context.Context, email string, passwordHash []byte) (int64, error) {
	return 0, nil
}

func (s *Storage) FindUser(ctx context.Context, email string) (models.User, error) {
	return models.User{}, nil
}

func (s *Storage) FindApp(ctx context.Context, appId int) (models.App, error) {
	return models.App{}, nil
}
