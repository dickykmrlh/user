package port

import (
	"github.com/dickykmrlh/user/internal/core/domain"
	"github.com/google/uuid"
)

type UserService interface {
	GetUser(id uuid.UUID) (user *domain.User, err error)
	Create(user *domain.User) (err error)
}

type UserRepository interface {
	GetUser(id uuid.UUID) (user *domain.User, err error)
	Save(user *domain.User) (err error)
}
