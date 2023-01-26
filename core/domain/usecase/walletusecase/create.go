package walletusecase

import (
	"strings"
	"time"

	"github.com/EliasSantiago/app-bank-transactions/core/dto"
	"github.com/google/uuid"
)

func (usecase usecase) Create(walletRequest *dto.CreateWalletRequest) (*dto.CreateWalletResponse, error) {
	uuid := uuid.New()
	store := &dto.CreateWalletStore{
		ID:        strings.Replace(uuid.String(), "-", "", -1),
		UserID:    walletRequest.UserID,
		Balance:   walletRequest.Balance,
		CreatedAt: time.Now().UnixNano() / int64(time.Millisecond),
		UpdatedAt: time.Now().UnixNano() / int64(time.Millisecond),
	}
	err := usecase.repository.Create(store)
	if err != nil {
		return nil, err
	}
	response := &dto.CreateWalletResponse{
		ID:        store.ID,
		Balance:   store.Balance,
		CreatedAt: store.CreatedAt,
		UpdatedAt: store.UpdatedAt,
	}
	return response, nil
}
