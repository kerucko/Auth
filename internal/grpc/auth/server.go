package auth

import (
	"context"

	api "github.com/kerucko/auth/pkg/api/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Auth interface {
	Register(ctx context.Context, email string, password string) (userId int64, err error)
	Login(ctx context.Context, email string, password string, appId int) (tokem string, err error)
}

type ServerAPI struct {
	api.UnimplementedAuthServer

	auth Auth
}

func Register(grpc *grpc.Server, auth Auth) {
	api.RegisterAuthServer(grpc, &ServerAPI{auth: auth})
}

func (s *ServerAPI) Register(ctx context.Context, request *api.RegisterRequest) (*api.RegisterResponse, error) {
	if request.Email == "" || request.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "empty email or password")
	}

	userId, err := s.auth.Register(ctx, request.GetEmail(), request.GetPassword())
	if err != nil {
		// TODO: существующий email
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &api.RegisterResponse{UserId: userId}, nil
}

func (s *ServerAPI) Login(ctx context.Context, request *api.LoginRequest) (*api.LoginResponse, error) {
	if request.Email == "" || request.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "empty email or password")
	}
	if request.AppId == 0 {
		return nil, status.Error(codes.InvalidArgument, "empty app id")
	}

	token, err := s.auth.Login(ctx, request.GetEmail(), request.GetPassword(), int(request.GetAppId()))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &api.LoginResponse{Token: token}, nil
}
