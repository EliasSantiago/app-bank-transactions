package di

import (
	"github.com/EliasSantiago/app-bank-transactions/adapter/http/walletservice"
	"github.com/EliasSantiago/app-bank-transactions/adapter/postgres"
	"github.com/EliasSantiago/app-bank-transactions/adapter/postgres/walletrepository"
	"github.com/EliasSantiago/app-bank-transactions/core/domain"
	"github.com/EliasSantiago/app-bank-transactions/core/domain/usecase/walletusecase"
)

func ConfigWalletDI(conn postgres.PoolInterface) domain.WalletService {
	walletRepository := walletrepository.New(conn)
	walletUseCase := walletusecase.New(walletRepository)
	walletService := walletservice.New(walletUseCase)
	return walletService
}
