package auth

import (
	"context"

	api "github.com/kerucko/auth/pkg/api/auth"
	"google.golang.org/grpc"
)

type ServerAPI struct {
	api.UnimplementedAuthServer
}

func Register(grpc *grpc.Server) {
	api.RegisterAuthServer(grpc, &ServerAPI{})
}

func (s *ServerAPI) Register(ctx context.Context, request *api.RegisterRequest) (*api.RegisterResponse, error) {
	return &api.RegisterResponse{}, nil
}

func (s *ServerAPI) Login(ctx context.Context, request *api.LoginRequest) (*api.LoginResponse, error) {
	return &api.LoginResponse{}, nil
}
