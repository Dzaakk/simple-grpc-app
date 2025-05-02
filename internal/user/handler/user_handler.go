package handler

import (
	"context"

	pb "github.com/dzaakk/simple-grpc-app/internal/user/generated"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}
func (h *UserHandler) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.StandardResponse, error) {
	return &pb.StandardResponse{Status: "ok", Message: "created"}, nil
}

func (h *UserHandler) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.StandardResponse, error) {
	return &pb.StandardResponse{Status: "ok", Message: "updated"}, nil
}

func (h *UserHandler) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.StandardResponse, error) {
	return &pb.StandardResponse{Status: "ok", Message: "deleted"}, nil
}

func (h *UserHandler) GetAllUsers(ctx context.Context, req *pb.Empty) (*pb.GetAllUsersResponse, error) {
	return &pb.GetAllUsersResponse{Status: "ok", Message: "fetched", Data: []*pb.User{}}, nil
}

func (h *UserHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.AuthResponse, error) {
	return &pb.AuthResponse{Status: "ok", Message: "logged in", Token: "dummy-token"}, nil
}

func (h *UserHandler) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.StandardResponse, error) {
	return &pb.StandardResponse{Status: "ok", Message: "logged out"}, nil
}
