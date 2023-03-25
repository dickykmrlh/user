package domain

import (
	"database/sql/driver"
	"fmt"
	"strings"

	user_v1 "github.com/dickykmrlh/protos/gen/go/user/v1"
	"github.com/google/uuid"
)

type UserRole int

const (
	UnknownUserRole UserRole = iota
	SuperAdminUserRole
	CustomerUserRole
)

func (r UserRole) String() string {
	switch r {
	case SuperAdminUserRole:
		return "super_admin"
	case CustomerUserRole:
		return "customer"
	default:
		return "unknown"
	}
}

func ToUserRole(s string) UserRole {
	switch strings.ToLower(s) {
	case "super_admin":
		return SuperAdminUserRole
	case "customer":
		return CustomerUserRole
	default:
		return UnknownUserRole
	}
}

func (u *UserRole) Value() (driver.Value, error) {
	return u.String(), nil
}

func (r *UserRole) Scan(value any) error {
	if value == nil {
		return nil
	}
	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("scan user role type assertion to byte failed")
	}

	*r = ToUserRole(string(b))

	return nil
}

type User struct {
	ID          uuid.UUID `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Role        UserRole  `json:"role"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
}

func (u User) Tov1User() *user_v1.User {
	return &user_v1.User{
		Id:          u.ID.String(),
		FirstName:   u.FirstName,
		LastName:    u.LastName,
		Role:        user_v1.UserRole(u.Role),
		PhoneNumber: u.PhoneNumber,
		Email:       u.Email,
	}
}

func NewUser(firstName, lastName, role, phoneNumber, email string) (*User, error) {
	if ToUserRole(role) == UnknownUserRole {
		return nil, fmt.Errorf("user role is unknown, role: %s", role)
	}
	return &User{
		ID:          uuid.New(),
		FirstName:   firstName,
		LastName:    lastName,
		Role:        ToUserRole(role),
		PhoneNumber: phoneNumber,
		Email:       email,
	}, nil
}
