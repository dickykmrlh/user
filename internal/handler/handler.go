package handler

import (
	"context"
	"errors"

	connect_go "github.com/bufbuild/connect-go"
	user_v1 "github.com/dickykmrlh/protos/gen/go/user/v1"
	port "github.com/dickykmrlh/user/internal/core/ports"
)

type UserHandler struct {
	userService port.UserService
}

func NewUserHandler(userService port.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) GetUser(ctx context.Context, req *connect_go.Request[user_v1.GetUserRequest]) (*connect_go.Response[user_v1.GetUserResponse], error) {
	return nil, errors.New("not implemented")
}

func (h *UserHandler) CreateUser(ctx context.Context, req *connect_go.Request[user_v1.CreateUserRequest]) (*connect_go.Response[user_v1.CreateUserResponse], error) {
	return nil, errors.New("not implemented")
}
