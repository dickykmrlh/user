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

func (s *service) Create(ctx context.Context, firstName, lastName, role, phoneNumber, email string) (err error) {
	user, err := domain.NewUser(firstName, lastName, role, phoneNumber, email)
	if err != nil {
		return err
	}

	return s.userRepo.Save(ctx, user)
}
