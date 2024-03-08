package grpcapp

import (
	"fmt"
	"log"
	"net"

	"github.com/kerucko/auth/internal/server"

	"google.golang.org/grpc"
)

type App struct {
	GrpcServer *grpc.Server
	Port       int
}

func NewApp(port int, auth server.Auth) *App {
	grpcServer := grpc.NewServer()
	server.Register(grpcServer, auth)
	return &App{
		GrpcServer: grpcServer,
		Port:       port,
	}
}

func (a *App) Run() error {
	op := "grpcapp.Run"
	log.Println("Starting gRPC server")
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.Port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Println("Running gRPC server on port", a.Port)
	if err := a.GrpcServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (a *App) Stop() {
	log.Println("Stopping gRPC server")
	a.GrpcServer.GracefulStop()
}
