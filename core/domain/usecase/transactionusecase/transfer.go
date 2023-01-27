package transactionusecase

import (
	"github.com/EliasSantiago/app-bank-transactions/core/dto"
)

func (usecase usecase) Transfer(transactionRequest *dto.CreateTransactionResponse) error {
	wallet, err := usecase.repository.Balance(transactionRequest.From)
	if err != nil {
		return err
	}
	if transactionRequest.Value > wallet.Balance {
		return err
	}
	store := &dto.CreateTransactionStore{
		ID:        transactionRequest.ID,
		From:      transactionRequest.From,
		To:        transactionRequest.To,
		Value:     transactionRequest.Value,
		Status:    false,
		CreatedAt: transactionRequest.CreatedAt,
	}
	err = usecase.repository.Transfer(store)
	if err != nil {
		return err
	}
	return nil
}
