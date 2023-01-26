package transactionrepository

import (
	"context"

	"github.com/EliasSantiago/app-bank-transactions/core/domain"
	"github.com/EliasSantiago/app-bank-transactions/core/dto"
)

func (repository repository) Create(transactionRequest *dto.CreateTransactionStore) error {
	ctx := context.Background()
	transaction := domain.Transaction{}
	err := repository.db.QueryRow(
		ctx,
		"INSERT INTO transactions (id, from, to, status, value, created_at) VALUES ($1, $2, $3, $4, $5, $6) returning *",
		transactionRequest.ID,
		transactionRequest.From,
		transactionRequest.To,
		transactionRequest.Status,
		transactionRequest.Value,
		transactionRequest.CreatedAt,
	).Scan(
		&transaction.ID,
		&transaction.From,
		&transaction.To,
		&transaction.Status,
		&transaction.Value,
		&transaction.CreatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}
