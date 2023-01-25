package userusecase

import (
	"strings"
	"time"

	"github.com/EliasSantiago/app-bank-transactions/core/dto"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (usecase usecase) Create(userRequest *dto.CreateUserRequest) (*dto.CreateUserResponse, error) {
	uuid := uuid.New()
	password, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), 12)
	if err != nil {
		return nil, err
	}
	store := &dto.CreateUserStore{
		ID:        strings.Replace(uuid.String(), "-", "", -1),
		Name:      userRequest.Name,
		Email:     userRequest.Email,
		Cpf:       userRequest.Cpf,
		Cnpj:      userRequest.Cnpj,
		Password:  password,
		CreatedAt: time.Now().UnixNano() / int64(time.Millisecond),
	}
	err = usecase.repository.Create(store)
	if err != nil {
		return nil, err
	}

	response := &dto.CreateUserResponse{
		ID:        store.ID,
		Name:      store.Name,
		Email:     store.Email,
		Cpf:       store.Cpf,
		Cnpj:      store.Cnpj,
		CreatedAt: store.CreatedAt,
	}
	return response, nil
}
