package port

import (
	"github.com/dickykmrlh/user/internal/core/domain"
	"github.com/docker/distribution/uuid"
)

type UserService interface {
	GetUser(id uuid.UUID) (user *domain.User)
	Create(user *domain.User) (err error)
}

type UserRepository interface {
	GetUser(id uuid.UUID) (user *domain.User)
	Save(user *domain.User) (err error)
}
