package domain

import (
	"net/http"

	"github.com/EliasSantiago/app-bank-transactions/core/dto"
)

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Cpf       string `json:"cpf"`
	Cnpj      string `json:"cnpj"`
	Password  string `json:"password"`
	CreatedAt int64  `json:"created_at"`
}

type UserService interface {
	Create(response http.ResponseWriter, request *http.Request)
}

type UserUseCase interface {
	Create(userRequest *dto.CreateUserRequest) (*dto.CreateUserResponse, error)
}

type UserRepository interface {
	Create(userRequest *dto.CreateUserStore) error
}
