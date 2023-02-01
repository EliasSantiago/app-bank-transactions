package walletrepository_test

import (
	"fmt"
	"testing"

	"github.com/EliasSantiago/app-bank-transactions/adapter/postgres/walletrepository"
	"github.com/EliasSantiago/app-bank-transactions/core/domain"
	"github.com/EliasSantiago/app-bank-transactions/core/dto"
	"github.com/bxcodec/faker/v3"
	"github.com/pashagolub/pgxmock"
	"github.com/stretchr/testify/require"
)

func setupCreate() ([]string, dto.CreateWalletStore, domain.Wallet, pgxmock.PgxPoolIface) {
	cols := []string{"id", "user_id", "balance", "create_at", "updated_at"}
	fakeWalletRequest := dto.CreateWalletStore{}
	fakeWalletDBResponse := domain.Wallet{}
	faker.FakeData(&fakeWalletRequest)
	faker.FakeData(&fakeWalletDBResponse)

	mock, _ := pgxmock.NewPool()
	return cols, fakeWalletRequest, fakeWalletDBResponse, mock
}

func TestCreate(t *testing.T) {
	cols, fakeWalletStore, fakeWalletDBResponse, mock := setupCreate()
	defer mock.Close()
	mock.ExpectQuery("INSERT INTO wallets (.+)").WithArgs(
		fakeWalletStore.ID,
		fakeWalletStore.UserID,
		fakeWalletStore.Balance,
		fakeWalletStore.CreatedAt,
		fakeWalletStore.UpdatedAt,
	).WillReturnRows(pgxmock.NewRows(cols).AddRow(
		fakeWalletDBResponse.ID,
		fakeWalletDBResponse.UserID,
		fakeWalletDBResponse.Balance,
		fakeWalletDBResponse.CreatedAt,
		fakeWalletDBResponse.UpdatedAt,
	))

	sut := walletrepository.New(mock)
	err := sut.Create(&fakeWalletStore)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
	require.Nil(t, err)
	// require.NotEmpty(t, wallet.ID)
	// require.Equal(t, wallet.UserID, fakeWalletDBResponse.UserID)
	// require.Equal(t, wallet.Balance, fakeWalletDBResponse.Balance)
	// require.Equal(t, wallet.CreatedAt, fakeWalletDBResponse.CreatedAt)
	// require.Equal(t, wallet.UpdatedAt, fakeWalletDBResponse.UpdatedAt)
}

func TestCreate_DBError(t *testing.T) {
	_, fakeWalletStore, _, mock := setupCreate()
	defer mock.Close()
	mock.ExpectQuery("INSERT INTO wallets (.+)").WithArgs(
		fakeWalletStore.ID,
		fakeWalletStore.UserID,
		fakeWalletStore.Balance,
		fakeWalletStore.CreatedAt,
		fakeWalletStore.UpdatedAt,
	).WillReturnError(fmt.Errorf("ANY DATABASE ERROR"))
	sut := walletrepository.New(mock)
	err := sut.Create(&fakeWalletStore)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
	require.NotNil(t, err)
	// require.Nil(t, wallet)
}
