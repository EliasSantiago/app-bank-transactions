package dto

import (
	"encoding/json"
	"io"
)

type CreateTransactionRequest struct {
	From  string  `json:"from"`
	To    string  `json:"to"`
	Value float64 `json:"value"`
}

type CreateTransactionResponse struct {
	ID        string  `json:"id"`
	From      string  `json:"from"`
	To        string  `json:"to"`
	Value     float64 `json:"value"`
	Status    bool    `json:"status"`
	CreatedAt int64   `json:"created_at"`
}

type CreateTransactionStore struct {
	ID        string  `db:"id"`
	From      string  `db:"from"`
	To        string  `db:"to"`
	Value     float64 `db:"value"`
	Status    bool    `db:"status"`
	CreatedAt int64   `db:"created_at"`
	UpdatedAt int64   `db:"updated_at"`
}

func FromJSONCreateTransactionRequest(body io.Reader) (*CreateTransactionRequest, error) {
	createTransactionRequest := CreateTransactionRequest{}
	if err := json.NewDecoder(body).Decode(&createTransactionRequest); err != nil {
		return nil, err
	}
	return &createTransactionRequest, nil
}
