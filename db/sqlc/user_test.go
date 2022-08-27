package db

import (
	"context"
	"testing"

	"github.com/alessandrorbianchi/gofinance-backend/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		StUsername: util.RandomString(6),
		StPassword: util.RandomString(12),
		StEmail:    util.RandomEmail(8),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.StUsername, user.StUsername)
	require.Equal(t, arg.StPassword, user.StPassword)
	require.Equal(t, arg.StEmail, user.StEmail)
	require.NotEmpty(t, user.DtCreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.StUsername)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.StUsername, user2.StUsername)
	require.Equal(t, user1.StPassword, user2.StPassword)
	require.Equal(t, user1.StEmail, user2.StEmail)
	require.NotEmpty(t, user2.DtCreatedAt)
}

func TestGetUserById(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUserById(context.Background(), user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.StUsername, user2.StUsername)
	require.Equal(t, user1.StPassword, user2.StPassword)
	require.Equal(t, user1.StEmail, user2.StEmail)
	require.NotEmpty(t, user2.DtCreatedAt)
}
