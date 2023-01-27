package transactionusecase

import (
	"github.com/EliasSantiago/app-bank-transactions/core/dto"
)

func (usecase usecase) Balance(id string) (*dto.GetBalanceResponse, error) {
	wallet, err := usecase.repository.Balance(id)
	if err != nil {
		return nil, err
	}
	response := &dto.GetBalanceResponse{
		Balance: wallet.Balance,
	}
	return response, nil
}
