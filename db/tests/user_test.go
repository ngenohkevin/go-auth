package tests

import (
	"context"
	db "github.com/ngenohkevin/go-auth/db/sqlc"
	"github.com/ngenohkevin/go-auth/utils"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomUser(t *testing.T) db.User {
	hashedPassword, err := utils.HashPassword(utils.RandomString(8))
	require.NoError(t, err)
	arg := db.CreateUserParams{
		Username:       utils.RandomUser(),
		Email:          utils.RandomEmail(),
		FullName:       utils.RandomUser(),
		HashedPassword: hashedPassword,
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)

	require.NotZero(t, user.CreatedAt)
	require.True(t, user.PasswordChangedAt.IsZero())

	return user
}
