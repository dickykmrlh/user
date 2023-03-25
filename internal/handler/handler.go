package handler

import (
	"context"
	"errors"

	"github.com/bufbuild/connect-go"
	connect_go "github.com/bufbuild/connect-go"
	user_v1 "github.com/dickykmrlh/protos/gen/go/user/v1"
	port "github.com/dickykmrlh/user/internal/core/ports"
	"github.com/google/uuid"
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
	err := req.Msg.Validate()
	if err != nil {
		return nil, err
	}

	user, err := h.userService.GetUser(ctx, uuid.MustParse(req.Msg.GetId()))
	if err != nil {
		return nil, err
	}

	if user == nil {
		return &connect_go.Response[user_v1.GetUserResponse]{}, nil
	}

	return connect.NewResponse(&user_v1.GetUserResponse{
		User: user.Tov1User(),
	}), nil
}

func (h *UserHandler) CreateUser(ctx context.Context, req *connect_go.Request[user_v1.CreateUserRequest]) (*connect_go.Response[user_v1.CreateUserResponse], error) {
	err := req.Msg.Validate()
	if err != nil {
		return nil, err
	}

	if req.Msg.GetUser().GetPhoneNumber() == "" && req.Msg.GetUser().GetEmail() == "" {
		return nil, errors.New("missing contact information")
	}

	return nil, errors.New("not implemented")
}
