package walletusecase

import "github.com/EliasSantiago/app-bank-transactions/core/domain"

type usecase struct {
	repository domain.WalletRepository
}

func New(repository domain.WalletRepository) domain.WalletUseCase {
	return &usecase{
		repository: repository,
	}
}
