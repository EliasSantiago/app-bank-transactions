package dto

import (
	"encoding/json"
	"io"
)

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required,name"`
	Email    string `json:"email" validate:"required,email"`
	Cpf      string `json:"cpf"`
	Cnpj     string `json:"cnpj"`
	Password string `json:"password" validate:"required,password"`
}

type CreateUserResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Cpf       string `json:"cpf"`
	Cnpj      string `json:"cnpj"`
	CreatedAt int64  `json:"created_at"`
}

type CreateUserStore struct {
	ID        string `db:"id"`
	Name      string `db:"name"`
	Email     string `db:"email"`
	Cpf       string `db:"cpf"`
	Cnpj      string `db:"cnpj"`
	Password  []byte `db:"password" json:"-"`
	CreatedAt int64  `db:"created_at"`
}

func FromJSONCreateUserRequest(body io.Reader) (*CreateUserRequest, error) {
	createUserRequest := CreateUserRequest{}
	if err := json.NewDecoder(body).Decode(&createUserRequest); err != nil {
		return nil, err
	}
	return &createUserRequest, nil
}
