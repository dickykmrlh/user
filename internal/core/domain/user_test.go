package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	t.Run("fail to create new user, role is unkown", func(t *testing.T) {
		_, err := NewUser("Sheldon", "Cooper", "62812345678", "", 4)
		assert.Equal(t, "user role is unknown, role: 4", err.Error())
	})

	t.Run("fail to create new user, must had contact information, phone number or email is a must", func(t *testing.T) {
		_, err := NewUser("Sheldon", "Cooper", "", "", 2)
		assert.Equal(t, "missing contact information", err.Error())
	})

	t.Run("fail to create new user, phone number is not matched", func(t *testing.T) {
		_, err := NewUser("Sheldon", "Cooper", "+xx812345678", "", 2)
		assert.Equal(t, "phone number didnt matched", err.Error())
	})

	t.Run("successfully create new user, phone number with +", func(t *testing.T) {
		user, err := NewUser("Sheldon", "Cooper", "+628123456781234", "", 2)
		assert.Nil(t, err)
		assert.Equal(t, "Sheldon", user.FirstName)
		assert.Equal(t, "Cooper", user.LastName)
		assert.Equal(t, CustomerUserRole, user.Role)
		assert.Equal(t, "+628123456781234", user.PhoneNumber)
	})

	t.Run("successfully create new user, phone number without +", func(t *testing.T) {
		user, err := NewUser("Sheldon", "Cooper", "62812345678", "", 2)
		assert.Nil(t, err)
		assert.Equal(t, "Sheldon", user.FirstName)
		assert.Equal(t, "Cooper", user.LastName)
		assert.Equal(t, CustomerUserRole, user.Role)
		assert.Equal(t, "+62812345678", user.PhoneNumber)
	})

	t.Run("successfully create new user, empty phone number but had email", func(t *testing.T) {
		user, err := NewUser("Sheldon", "Cooper", "", "sheldon@gmail.com", 2)
		assert.Nil(t, err)
		assert.Equal(t, "Sheldon", user.FirstName)
		assert.Equal(t, "Cooper", user.LastName)
		assert.Equal(t, CustomerUserRole, user.Role)
		assert.Equal(t, "sheldon@gmail.com", user.Email)
	})
}
