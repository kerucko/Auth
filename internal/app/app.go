package app

import (
	grpcapp "github.com/kerucko/auth/internal/app/grpc"
	"github.com/kerucko/auth/internal/config"
)

type App struct {
	GrpcServer *grpcapp.Server
}

func NewApp(cfg config.Config) *App {
	grpcApp := grpcapp.NewServer(cfg.Grpc.Port)

	return &App{
		GrpcServer: grpcApp,
	}
}
