package domain

import (
	"net/http"

	"github.com/EliasSantiago/app-bank-transactions/core/dto"
)

type Transaction struct {
	ID        string  `json:"id"`
	From      string  `json:"from"`
	To        string  `json:"to"`
	Value     float64 `json:"value"`
	Status    bool    `json:"status"`
	CreatedAt int64   `json:"created_at"`
}

type TransactionService interface {
	Create(response http.ResponseWriter, request *http.Request)
	Consumer()
}

type TransactionUseCase interface {
	Create(transactionRequest *dto.CreateTransactionRequest) (*dto.CreateTransactionResponse, error)
	Transfer(transactionRequest *dto.CreateTransactionResponse) error
	Balance(id string) (*dto.GetBalanceResponse, error)
	Consumer()
	Check(id string) (bool, error)
}

type TransactionRepository interface {
	Transfer(transactionRequest *dto.CreateTransactionStore) error
	Balance(id string) (*dto.GetBalanceStore, error)
	UpdateBalance(value float64, id string) error
	Check(id string) (bool, error)
}
