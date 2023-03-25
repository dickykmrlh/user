package service

import (
	"context"

	"github.com/dickykmrlh/user/internal/core/domain"
	port "github.com/dickykmrlh/user/internal/core/ports"
	"github.com/google/uuid"
)

type service struct {
	userRepo port.UserRepository
}

func NewUserService(userRepo port.UserRepository) *service {
	return &service{
		userRepo: userRepo,
	}
}

func (s *service) GetUser(ctx context.Context, id uuid.UUID) (user *domain.User, err error) {
	return s.userRepo.GetUser(ctx, id)
}

func (s *service) CreateUser(ctx context.Context, firstName, lastName, phoneNumber, email string, role int) (user *domain.User, err error) {
	user, err = domain.NewUser(firstName, lastName, phoneNumber, email, role)
	if err != nil {
		return nil, err
	}

	err = s.userRepo.Save(ctx, user)
	return
}
