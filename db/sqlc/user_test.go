package db

import (
	"context"
	"testing"

	"github.com/john9101/simplebank/util"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) User {
	hashedPassword, err := util.HashPassword("secret")
	require.NoError(t, err)

	arg := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}

	ctx := context.Background()

	user, err := testQueries.CreateUser(ctx, arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)

	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user
}
