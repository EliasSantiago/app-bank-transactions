package transactionusecase

import (
	"github.com/EliasSantiago/app-bank-transactions/core/dto"
)

func (usecase usecase) Create(transactionRequest *dto.CreateTransactionRequest) (*dto.CreateTransactionResponse, error) {
	//uuid := uuid.New()
	store := &dto.CreateTransactionStore{}
	err := usecase.repository.Create(store)
	if err != nil {
		return nil, err
	}
	response := &dto.CreateTransactionResponse{}
	return response, nil
}
