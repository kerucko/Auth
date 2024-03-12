package storage

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/kerucko/auth/internal/models"
)

var (
	ErrUserNotFound = fmt.Errorf("user not found")
	ErrUserExists   = fmt.Errorf("user already exists")
	ErrAppNotFound  = fmt.Errorf("app not found")
)

type Storage struct {
	db *pgx.Conn
}

func New(dbPath string, timeout time.Duration) *Storage {
	op := "storage.NewStorage"

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	db, err := pgx.Connect(ctx, dbPath)
	if err != nil {
		log.Printf("%s: %s", op, err.Error())
	}
	log.Println("Connected to db")
	return &Storage{db: db}

}

func (s *Storage) SaveUser(ctx context.Context, email string, passwordHash []byte) (int64, error) {
	op := "storage.SaveUser"

	request := `
		INSERT INTO users (email, password_hash)
		VALUES ($1, $2)
		RETURNING id;
	`

	var userId int64
	err := s.db.QueryRow(ctx, request, email, passwordHash).Scan(&userId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			log.Printf("%s: %s", op, ErrUserExists.Error())
			return 0, fmt.Errorf("%s: %w", op, ErrUserExists)
		}

		log.Printf("%s: %s", op, err.Error())
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return userId, nil
}

func (s *Storage) FindUser(ctx context.Context, email string) (models.User, error) {
	op := "storage.FindUser"

	request := `
		SELECT id, email, password_hash
		FROM users
		WHERE email = $1;
	`

	var user models.User
	err := s.db.QueryRow(ctx, request, email).Scan(&user.Id, &user.Email, &user.PasswordHash)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			log.Printf("%s: %s", op, ErrUserNotFound.Error())
			return models.User{}, fmt.Errorf("%s: %w", op, ErrUserNotFound)
		}

		log.Printf("%s: %s", op, err.Error())
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}

func (s *Storage) FindApp(ctx context.Context, appId int) (models.App, error) {
	op := "storage.FindApp"

	request := `
		SELECT id, secret
		FROM apps
		WHERE id = $1;
	`

	var app models.App
	err := s.db.QueryRow(ctx, request, appId).Scan(&app.Id, &app.Secret)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			log.Printf("%s: %s", op, ErrAppNotFound.Error())
			return models.App{}, fmt.Errorf("%s: %w", op, ErrAppNotFound)
		}

		log.Printf("%s: %s", op, err.Error())
		return models.App{}, fmt.Errorf("%s: %w", op, err)
	}

	return app, nil
}
