package app

import (
	"fmt"

	"github.com/kerucko/auth/internal/app/grpcapp"
	"github.com/kerucko/auth/internal/config"
	"github.com/kerucko/auth/internal/services/auth"
	"github.com/kerucko/auth/internal/storage"
)

type App struct {
	GrpcServer *grpcapp.App
}

func NewApp(cfg config.Config) *App {
	dbPath := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.Dbname)
	storage := storage.New(dbPath, cfg.Database.Timeout)
	authService := auth.NewAuth(storage, storage, storage, cfg.ToketExpiration)
	grpcApp := grpcapp.NewApp(cfg.Grpc.Port, authService)

	return &App{
		GrpcServer: grpcApp,
	}
}