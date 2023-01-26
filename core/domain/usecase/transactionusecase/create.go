package transactionusecase

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/EliasSantiago/app-bank-transactions/core/dto"
	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

func Publish(ch *amqp.Channel, transactionRequest *dto.CreateTransactionRequest) error {
	body, err := json.Marshal(transactionRequest)
	if err != nil {
		return err
	}
	err = ch.Publish(
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
	Publish(ch, transactionRequest)
	return response, nil
}
