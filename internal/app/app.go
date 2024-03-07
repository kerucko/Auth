package app

import (
	"github.com/kerucko/auth/internal/app/grpcapp"
	"github.com/kerucko/auth/internal/config"
	"github.com/kerucko/auth/internal/services/auth"
	"github.com/kerucko/auth/internal/storage"
)

type App struct {
	GrpcServer *grpcapp.Server
}

func NewApp(cfg config.Config) *App {
	storage := storage.NewStorage(cfg.Database.Host, cfg.Database.Port, cfg.Database.Dbname, cfg.Database.User, cfg.Database.Password, cfg.Database.Timeout)
	authService := auth.NewAuth(storage, storage, storage, cfg.ToketExpiration)
	grpcApp := grpcapp.NewServer(cfg.Grpc.Port, authService)

	return &App{
		GrpcServer: grpcApp,
	}
}
