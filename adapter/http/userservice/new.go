package userservice

import (
	"github.com/EliasSantiago/app-bank-transactions/core/domain"
)

type service struct {
	usecase domain.UserUseCase
}

func New(usecase domain.UserUseCase) domain.UserService {
	return &service{
		usecase: usecase,
	}
}
