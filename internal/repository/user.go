package repository

import (
	"database/sql"

	"github.com/dickykmrlh/user/internal/core/domain"
	"github.com/docker/distribution/uuid"
)

type UserRepo struct {
	db *sql.DB
}

func New(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (u *UserRepo) GetUser(id uuid.UUID) (user *domain.User) {
	return nil
}

func (u *UserRepo) Save(user *domain.User) (err error) {
	return nil
}
