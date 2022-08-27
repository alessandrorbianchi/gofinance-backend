package db

import (
	"context"
	"testing"
	"time"

	"github.com/alessandrorbianchi/gofinance-backend/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	category := createRandomCategory(t)
	arg := CreateAccountParams{
		CoUserID:      category.CoUserID,
		CoCategoryID:  category.ID,
		StTitle:       util.RandomString(12),
		StType:        category.StType,
		StDescription: util.RandomString(20),
		VlValue:       10,
		DtDate:        time.Now(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.CoUserID, account.CoUserID)
	require.Equal(t, arg.CoCategoryID, account.CoCategoryID)
	require.Equal(t, arg.VlValue, account.VlValue)
	require.Equal(t, arg.StTitle, account.StTitle)
	require.Equal(t, arg.StType, account.StType)
	require.Equal(t, arg.StDescription, account.StDescription)

	require.NotEmpty(t, account.DtCreatedAt)
	require.NotEmpty(t, account.DtDate)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.CoUserID, account2.CoUserID)
	require.Equal(t, account1.CoCategoryID, account2.CoCategoryID)
	require.Equal(t, account1.VlValue, account2.VlValue)
	require.Equal(t, account1.StTitle, account2.StTitle)
	require.Equal(t, account1.StType, account2.StType)
	require.Equal(t, account1.StDescription, account2.StDescription)

	require.NotEmpty(t, account2.DtDate)
	require.NotEmpty(t, account2.DtCreatedAt)
}

func TestDeleteAccount(t *testing.T) {
	account := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)
}

func TestUpdateAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:            account1.ID,
		StTitle:       util.RandomString(12),
		StDescription: util.RandomString(20),
		VlValue:       15,
	}

	account2, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, arg.StTitle, account2.StTitle)
	require.Equal(t, arg.StDescription, account2.StDescription)
	require.Equal(t, arg.VlValue, account2.VlValue)
	require.Equal(t, account1.DtCreatedAt, account2.DtCreatedAt)
}

func TestListAccounts(t *testing.T) {
	lastAccount := createRandomAccount(t)

	arg := GetAccountsParams{
		CoUserID:      lastAccount.CoUserID,
		StType:        lastAccount.StType,
		CoCategoryID:  lastAccount.CoCategoryID,
		DtDate:        lastAccount.DtDate,
		StTitle:       lastAccount.StTitle,
		StDescription: lastAccount.StDescription,
	}

	accounts, err := testQueries.GetAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)

	for _, account := range accounts {
		require.Equal(t, lastAccount.ID, account.ID)
		require.Equal(t, lastAccount.CoUserID, account.CoUserID)
		require.Equal(t, lastAccount.StTitle, account.StTitle)
		require.Equal(t, lastAccount.StDescription, account.StDescription)
		require.Equal(t, lastAccount.VlValue, account.VlValue)
		require.NotEmpty(t, lastAccount.DtCreatedAt)
		require.NotEmpty(t, lastAccount.DtDate)
	}
}

func TestListGetReports(t *testing.T) {
	lastAccount := createRandomAccount(t)

	arg := GetAccountsReportsParams{
		CoUserID: lastAccount.CoUserID,
		StType:   lastAccount.StType,
	}

	accounts, err := testQueries.GetAccountsReports(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)
}

func TestListGetGraph(t *testing.T) {
	lastAccount := createRandomAccount(t)

	arg := GetAccountsGraphParams{
		CoUserID: lastAccount.CoUserID,
		StType:   lastAccount.StType,
	}

	graphValue, err := testQueries.GetAccountsGraph(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, graphValue)
}
