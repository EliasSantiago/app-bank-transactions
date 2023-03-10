package transactionusecase

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/EliasSantiago/app-bank-transactions/core/domain/usecase/errors"
	"github.com/EliasSantiago/app-bank-transactions/core/dto"
	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

func Publish(ch *amqp.Channel, transactionRequest *dto.CreateTransactionResponse) error {
	body, err := json.Marshal(transactionRequest)
	if err != nil {
		return err
	}
	err = ch.PublishWithContext(context.Background(),
		"amq.direct",
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (usecase usecase) Create(transactionRequest *dto.CreateTransactionRequest) (*dto.CreateTransactionResponse, error) {
	uuid := uuid.New()
	check, err := usecase.Check(transactionRequest.From)
	if err != nil {
		return nil, errors.CheckDataExists()
	}
	if !check {
		return nil, errors.WalletNotExist()
	}
	if transactionRequest.Value <= 0 {
		return nil, errors.InvalidParameters()
	}
	wallet, err := usecase.repository.Balance(transactionRequest.From)
	if err != nil {
		return nil, errors.GetBalance()
	}
	if transactionRequest.Value > wallet.Balance {
		return nil, errors.InsufficientFunds()
	}
	store := &dto.CreateTransactionStore{
		ID:        strings.Replace(uuid.String(), "-", "", -1),
		From:      transactionRequest.From,
		To:        transactionRequest.To,
		Value:     transactionRequest.Value,
		Status:    false,
		CreatedAt: time.Now().UnixNano() / int64(time.Millisecond),
	}
	response := &dto.CreateTransactionResponse{
		ID:        store.ID,
		From:      store.From,
		To:        store.To,
		Value:     store.Value,
		Status:    "Pendente",
		CreatedAt: store.CreatedAt,
	}
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	defer ch.Close()
	Publish(ch, response)
	return response, nil
}
