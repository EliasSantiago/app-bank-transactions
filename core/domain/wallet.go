package domain

import (
	"net/http"

	"github.com/EliasSantiago/app-bank-transactions/core/dto"
)

type Wallet struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Balance   string `json:"balance"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type Balance struct {
	Balance string `json:"balance"`
}

type WalletService interface {
	Create(response http.ResponseWriter, request *http.Request)
}

type WalletUseCase interface {
	Create(walletRequest *dto.CreateWalletRequest) (*dto.CreateWalletResponse, error)
}

type WalletRepository interface {
	Create(walletRequest *dto.CreateWalletStore) error
}
