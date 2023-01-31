package di

import (
	"github.com/EliasSantiago/app-bank-transactions/adapter/http/transactionservice"
	"github.com/EliasSantiago/app-bank-transactions/adapter/postgres"
	"github.com/EliasSantiago/app-bank-transactions/adapter/postgres/transactionrepository"
	"github.com/EliasSantiago/app-bank-transactions/core/domain"
	"github.com/EliasSantiago/app-bank-transactions/core/domain/usecase/transactionusecase"
)

func ConfigTransactionQueueDI(conn postgres.PoolInterface) domain.TransactionService {
	transactionQueueRepository := transactionrepository.New(conn)
	transactionQueueUseCase := transactionusecase.New(transactionQueueRepository)
	transactionQueueService := transactionservice.New(transactionQueueUseCase)
	return transactionQueueService
}
