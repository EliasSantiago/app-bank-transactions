package walletusecase_test

import (
	"testing"

	"github.com/EliasSantiago/app-bank-transactions/core/domain"
	"github.com/EliasSantiago/app-bank-transactions/core/domain/mocks"
	"github.com/EliasSantiago/app-bank-transactions/core/domain/usecase/walletusecase"
	"github.com/EliasSantiago/app-bank-transactions/core/dto"
	"github.com/bxcodec/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	fakeRequestWallet := dto.CreateWalletRequest{}
	fakeDBWallet := domain.Wallet{}
	faker.FakeData(&fakeRequestWallet)
	faker.FakeData(&fakeDBWallet)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockWalletRepository := mocks.NewMockWalletRepository(mockCtrl)
	mockWalletRepository.EXPECT().Create(&fakeRequestWallet).Return(&fakeDBWallet, nil)

	sut := walletusecase.New(mockWalletRepository)
	wallet, err := sut.Create(&fakeRequestWallet)

	require.Nil(t, err)
	require.NotEmpty(t, wallet.ID)
	require.Equal(t, wallet.Balance, fakeDBWallet.Balance)
	require.Equal(t, wallet.CreatedAt, fakeDBWallet.CreatedAt)
}
