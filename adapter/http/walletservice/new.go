package walletservice

import "github.com/EliasSantiago/app-bank-transactions/core/domain"

type service struct {
	usecase domain.WalletUseCase
}

func New(usecase domain.WalletUseCase) domain.WalletService {
	return &service{
		usecase: usecase,
	}
}
