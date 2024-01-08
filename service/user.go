package service

import (
	"context"

	go_grpc_gauth_pg "github.com/Reigenleif/go-grpc-gauth-pg/api"
)

type UserService struct {
	go_grpc_gauth_pg.UnimplementedUserServer
}

func (s *UserService) GetUser(ctx context.Context, req *go_grpc_gauth_pg.GetUserRequest) (*go_grpc_gauth_pg.UserResponse, error) {
	

	return &go_grpc_gauth_pg.UserResponse{
		Id: "ax",
	}, nil
}

func (s *UserService) CreateUser(ctx context.Context, req *go_grpc_gauth_pg.CreateUserRequest) (*go_grpc_gauth_pg.UserResponse, error) {
	return &go_grpc_gauth_pg.UserResponse{
		Id: "ax",
	}, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *go_grpc_gauth_pg.UpdateUserRequest) (*go_grpc_gauth_pg.UserResponse, error) {
	return &go_grpc_gauth_pg.UserResponse{
		Id: "ax",
	}, nil
}

func (s *UserService) DeleteUser(ctx context.Context, req *go_grpc_gauth_pg.DeleteUserRequest) (*go_grpc_gauth_pg.UserResponse, error) {
	return &go_grpc_gauth_pg.UserResponse{
		Id: "ax",
	}, nil
}