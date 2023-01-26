package transactionservice

import "github.com/EliasSantiago/app-bank-transactions/core/domain"

type service struct {
	usecase domain.TransactionUseCase
}

func New(usecase domain.TransactionUseCase) domain.TransactionService {
	return &service{
		usecase: usecase,
	}
}
