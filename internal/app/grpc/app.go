package grpcapp

import (
	"fmt"
	"log"
	"net"

	authgrpc "github.com/kerucko/auth/internal/grpc/auth"
	"google.golang.org/grpc"
)

type Server struct {
	GrpcServer *grpc.Server
	Port       int
}

func NewServer(port int) *Server {
	grpcServer := grpc.NewServer()
	authgrpc.Register(grpcServer)
	return &Server{
		GrpcServer: grpcServer,
		Port:       port,
	}
}

func (a *Server) Run() error {
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

func (a *Server) Stop() {
	log.Println("Stopping gRPC server")
	a.GrpcServer.GracefulStop()
}