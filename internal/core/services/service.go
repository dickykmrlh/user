package service

import (
	"github.com/dickykmrlh/user/internal/core/domain"
	"github.com/docker/distribution/uuid"
)

type service struct {
}

func New() *service {
	return &service{}
}

func (s *service) GetUser(id uuid.UUID) (user *domain.User) {
	return nil
}

func (s *service) Create(user *domain.User) (err error) {
	return
}
