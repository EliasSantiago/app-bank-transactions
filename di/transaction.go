package di

import (
	"github.com/EliasSantiago/app-bank-transactions/adapter/http/transactionservice"
	"github.com/EliasSantiago/app-bank-transactions/adapter/postgres"
	"github.com/EliasSantiago/app-bank-transactions/adapter/postgres/transactionrepository"
	"github.com/EliasSantiago/app-bank-transactions/core/domain"
	"github.com/EliasSantiago/app-bank-transactions/core/domain/usecase/transactionusecase"
)

func ConfigTransactionDI(conn postgres.PoolInterface) domain.TransactionService {
	transactionRepository := transactionrepository.New(conn)
	transactionUseCase := transactionusecase.New(transactionRepository)
	transactionService := transactionservice.New(transactionUseCase)
	return transactionService
}
