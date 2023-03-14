package port

import (
	"context"

	"github.com/dickykmrlh/user/internal/core/domain"
	"github.com/google/uuid"
)

type UserService interface {
	GetUser(id uuid.UUID) (user *domain.User, err error)
	Create(user *domain.User) (err error)
}

type UserRepository interface {
	GetUser(ctx context.Context, id uuid.UUID) (user *domain.User, err error)
	Save(ctx context.Context, user *domain.User) (err error)
}
