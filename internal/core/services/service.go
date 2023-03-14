package service

import (
	"github.com/dickykmrlh/user/internal/core/domain"
	"github.com/google/uuid"
)

type service struct {
}

func New() *service {
	return &service{}
}

func (s *service) GetUser(id uuid.UUID) (user *domain.User, err error) {
	return nil, nil
}

func (s *service) Create(user *domain.User) (err error) {
	return
}
