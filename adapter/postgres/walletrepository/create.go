package walletrepository

import (
	"context"

	"github.com/EliasSantiago/app-bank-transactions/core/domain"
	"github.com/EliasSantiago/app-bank-transactions/core/dto"
)

func (repository repository) Create(walletRequest *dto.CreateWalletStore) error {
	ctx := context.Background()
	wallet := domain.Wallet{}
	err := repository.db.QueryRow(
		ctx,
		"INSERT INTO wallets (id, user_id, balance, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) returning *",
		walletRequest.ID,
		walletRequest.UserID,
		walletRequest.Balance,
		walletRequest.CreatedAt,
		walletRequest.UpdatedAt,
	).Scan(
		&wallet.ID,
		&wallet.UserID,
		&wallet.Balance,
		&wallet.CreatedAt,
		&wallet.UpdatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}
