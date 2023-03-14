package repository

import (
	"context"
	"fmt"
	"testing"

	"github.com/dickykmrlh/user/internal/core/domain"
	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func loadTestFixures(t *testing.T, fixtureName string) {
	t.Helper()
	fixtures, err := testfixtures.New(
		testfixtures.Database(GetDB()),
		testfixtures.Dialect("postgres"),
		testfixtures.UseAlterConstraint(),
		testfixtures.Directory(fmt.Sprintf("../../test/fixtures/%s", fixtureName)),
	)

	if err != nil {
		require.NoError(t, err)
	}

	err = fixtures.EnsureTestDatabase()
	if err != nil {
		require.NoError(t, err)
	}

	if err = fixtures.Load(); err != nil {
		require.NoError(t, err)
	}
}

func TestUserRepo_GetUser(t *testing.T) {
	err := godotenv.Overload("../../test/test.env")
	require.NoError(t, err)

	db := GetDB()
	u := &UserRepo{
		db: db,
	}

	loadTestFixures(t, "user")
	ctx := context.Background()

	t.Run("Get user", func(t *testing.T) {
		expctedUser := &domain.User{
			ID:          uuid.MustParse("a6671da6-90db-11ed-a9f2-170300000000"),
			FirstName:   "Dicky",
			LastName:    "Kamaulah",
			Role:        domain.SuperAdminUserRole,
			PhoneNumber: "+628123456789",
			Email:       "dicky@aja.com",
		}

		actualUser, err := u.GetUser(ctx, expctedUser.ID)
		assert.Nil(t, err)
		assert.Equal(t, expctedUser, actualUser)
	})
}
