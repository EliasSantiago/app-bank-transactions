package di

import (
	"github.com/EliasSantiago/app-bank-transactions/adapter/http/userservice"
	"github.com/EliasSantiago/app-bank-transactions/adapter/postgres"
	"github.com/EliasSantiago/app-bank-transactions/adapter/postgres/userrepository"
	"github.com/EliasSantiago/app-bank-transactions/core/domain"
	"github.com/EliasSantiago/app-bank-transactions/core/domain/usecase/userusecase"
)

func ConfigUserDI(conn postgres.PoolInterface) domain.UserService {
	userRepository := userrepository.New(conn)
	userUseCase := userusecase.New(userRepository)
	userService := userservice.New(userUseCase)
	return userService
}
