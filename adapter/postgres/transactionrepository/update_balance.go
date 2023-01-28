package transactionrepository

import (
	"context"

	"github.com/EliasSantiago/app-bank-transactions/core/domain"
)

func (repository repository) UpdateBalance(value float64, id string) error {
	wallet := domain.Wallet{}
	ctx := context.Background()
	err := repository.db.QueryRow(
		ctx,
		"UPDATE wallets SET balance = $1 WHERE id = $2 returning *",
		value,
		id,
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
