package userusecase

import (
	"github.com/EliasSantiago/app-bank-transactions/core/domain"
	"github.com/EliasSantiago/app-bank-transactions/core/dto"
)

func (usecase usecase) Create(userRequest *dto.CreateUserRequest) (*domain.User, error) {
	user, err := usecase.repository.Create(userRequest)
	if err != nil {
		return nil, err
	}
	return user, nil
}
