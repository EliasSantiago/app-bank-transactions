package dto

import (
	"encoding/json"
	"io"
)

type CreateWalletRequest struct {
	UserID  string  `json:"user_id"`
	Balance float64 `json:"balance"`
}

type CreateWalletResponse struct {
	ID        string  `json:"id"`
	Balance   float64 `json:"balance"`
	CreatedAt int64   `json:"created_at"`
	UpdatedAt int64   `json:"updated_at"`
}

type CreateWalletStore struct {
	ID        string  `db:"id"`
	UserID    string  `db:"user_id"`
	Balance   float64 `db:"balance"`
	CreatedAt int64   `db:"created_at"`
	UpdatedAt int64   `db:"updated_at"`
}

type GetBalanceResponse struct {
	Balance float64 `json:"balance"`
}

func FromJSONCreateWalletRequest(body io.Reader) (*CreateWalletRequest, error) {
	createWalletRequest := CreateWalletRequest{}
	if err := json.NewDecoder(body).Decode(&createWalletRequest); err != nil {
		return nil, err
	}
	return &createWalletRequest, nil
}
