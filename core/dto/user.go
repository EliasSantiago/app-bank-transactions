package dto

import (
	"encoding/json"
	"io"
)

type CreateUserRequest struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Cpf       string `json:"cpj"`
	Cnpj      string `json:"cnpj"`
	Password  string `json:"password"`
	CreatedAt int64  `json:"created_at"`
}

func FromJSONCreateUserRequest(body io.Reader) (*CreateUserRequest, error) {
	createUserRequest := CreateUserRequest{}

	if err := json.NewDecoder(body).Decode(&createUserRequest); err != nil {
		return nil, err
	}
	return &createUserRequest, nil
}
