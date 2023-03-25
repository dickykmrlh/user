package domain

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"regexp"
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

var _indonesianPhoneNumberPattern = regexp.MustCompile("^[+62]{2}[0-9]{9,}")

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

func NewUser(firstName, lastName, phoneNumber, email string, role int) (*User, error) {
	if UserRole(role).String() == UnknownUserRole.String() {
		return nil, fmt.Errorf("user role is unknown, role: %d", role)
	}

	if phoneNumber == "" && email == "" {
		return nil, errors.New("missing contact information")
	}

	if phoneNumber != "" {
		if matched := _indonesianPhoneNumberPattern.MatchString(phoneNumber); !matched {
			return nil, errors.New("phone number didnt matched")
		}
	}

	return &User{
		ID:        uuid.New(),
		FirstName: firstName,
		LastName:  lastName,
		Role:      UserRole(role),
		PhoneNumber: func(phone string) string {
			if phone == "" {
				return ""
			}

			if phone[0] != '+' {
				return "+" + phone
			}

			return phone
		}(phoneNumber),
		Email: email,
	}, nil
}
