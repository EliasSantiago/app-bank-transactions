package userrepository

import (
	"context"

	"github.com/EliasSantiago/app-bank-transactions/core/domain"
	"github.com/EliasSantiago/app-bank-transactions/core/dto"
)

func (repository repository) Create(userRequest *dto.CreateUserRequest) (*domain.User, error) {
	ctx := context.Background()
	user := domain.User{}
	err := repository.db.QueryRow(
		ctx,
		"INSERT INTO users (id, name, email, cpf, cnpj, password, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7) returning *",
		userRequest.ID,
		userRequest.Name,
		userRequest.Email,
		userRequest.Cpf,
		userRequest.Cnpj,
		userRequest.Password,
		userRequest.CreatedAt,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Cpf,
		&user.Cnpj,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
