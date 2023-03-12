package port

import (
	"github.com/dickykmrlh/user/internal/core/domain"
	"github.com/docker/distribution/uuid"
)

type UserService interface {
	GetUser(id uuid.UUID) (user *domain.User)
	Create(user *domain.User) (err error)
}
