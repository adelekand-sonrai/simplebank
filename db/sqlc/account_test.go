package db

import (
	"context"
	"database/sql"
	"simplebank/db/utils"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    utils.RandomOwner(),
		Balance:  utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	// create an account
	existingAccount := createRandomAccount(t)
	fetchedAccount, err := testQueries.GetAccount(context.Background(), existingAccount.ID)
	require.NoError(t, err)
	require.NotEmpty(t, fetchedAccount)

	require.Equal(t, existingAccount.Owner, fetchedAccount.Owner)
	require.Equal(t, existingAccount.Balance, fetchedAccount.Balance)
	require.Equal(t, existingAccount.Currency, fetchedAccount.Currency)
	require.Equal(t, existingAccount.ID, fetchedAccount.ID)
	require.Equal(t, existingAccount.CreatedAt, fetchedAccount.CreatedAt)
}

func TestUpdateAccount(t *testing.T) {
	existingAccount := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      existingAccount.ID,
		Balance: utils.RandomMoney(),
	}

	updatedAccount, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedAccount)

	require.Equal(t, existingAccount.Owner, updatedAccount.Owner)
	require.NotEqual(t, existingAccount.Balance, updatedAccount.Balance)
	require.Equal(t, existingAccount.Currency, updatedAccount.Currency)
	require.Equal(t, existingAccount.ID, updatedAccount.ID)
	require.Equal(t, existingAccount.CreatedAt, updatedAccount.CreatedAt)
}

func TestDeleteAccount(t *testing.T) {
	account := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)

	checkAccount, err2 := testQueries.GetAccount(context.Background(), account.ID)
	require.Error(t, err2)
	require.EqualError(t, err2, sql.ErrNoRows.Error())
	require.Empty(t, checkAccount)
}

func TestListAccount(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
