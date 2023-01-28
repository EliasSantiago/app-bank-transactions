package transactionusecase

import (
	"github.com/EliasSantiago/app-bank-transactions/core/domain/usecase/errors"
	"github.com/EliasSantiago/app-bank-transactions/core/dto"
)

func (usecase usecase) Transfer(transactionRequest *dto.CreateTransactionResponse) error {
	check, err := usecase.Check(transactionRequest.From)
	if err != nil {
		return errors.CheckDataExists()
	}
	if !check {
		return errors.WalletNotExist()
	}
	if transactionRequest.Value <= 0 {
		return errors.InvalidParameters()
	}
	balanceWalletFrom, err := usecase.repository.Balance(transactionRequest.From)
	if err != nil {
		return errors.GetBalance()
	}
	balanceWalletTo, err := usecase.repository.Balance(transactionRequest.To)
	if err != nil {
		return errors.GetBalance()
	}
	if transactionRequest.Value > balanceWalletFrom.Balance {
		return errors.InsufficientFunds()
	}
	fromBalance := balanceWalletFrom.Balance - transactionRequest.Value
	toBalance := balanceWalletTo.Balance + transactionRequest.Value
	err = usecase.repository.UpdateBalance(fromBalance, transactionRequest.From)
	if err != nil {
		return errors.UpdateBalance()
	}
	err = usecase.repository.UpdateBalance(toBalance, transactionRequest.To)
	if err != nil {
		return errors.UpdateBalance()
	}
	status := false
	if transactionRequest.Status == "Pendente" {
		status = false
	} else {
		status = true
	}
	store := &dto.CreateTransactionStore{
		ID:        transactionRequest.ID,
		From:      transactionRequest.From,
		To:        transactionRequest.To,
		Value:     transactionRequest.Value,
		Status:    status,
		CreatedAt: transactionRequest.CreatedAt,
	}
	err = usecase.repository.Transfer(store)
	if err != nil {
		return err
	}
	return nil
}
