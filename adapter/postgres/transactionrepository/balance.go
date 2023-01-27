package transactionrepository

import (
	"context"

	"github.com/EliasSantiago/app-bank-transactions/core/dto"
)

func (repository repository) Balance(id string) (*dto.GetBalanceStore, error) {
	ctx := context.Background()
	balance := &dto.GetBalanceStore{}
	err := repository.db.QueryRow(
		ctx,
		"SELECT balance FROM wallets WHERE id=$1",
		id,
	).Scan(
		&balance.Balance,
	)
	if err != nil {
		return nil, err
	}
	return balance, nil
}
