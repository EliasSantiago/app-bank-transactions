package transactionusecase

import (
	"encoding/json"
	"fmt"

	"github.com/EliasSantiago/app-bank-transactions/core/dto"
	"github.com/EliasSantiago/app-bank-transactions/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func (usecase usecase) Consumer() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()
	out := make(chan amqp.Delivery)
	go rabbitmq.Consume(ch, out)
	numWorkers := 2
	for i := 1; i <= numWorkers; i++ {
		for msg := range out {
			inputDTO := &dto.CreateTransactionResponse{}
			if err := json.Unmarshal(msg.Body, &inputDTO); err != nil {
				panic(err)
			}
			err = usecase.Transfer(inputDTO)
			if err != nil {
				// TODO: TRATAR PARA PUBLICAR EM UMA FILA DE ERROS
				panic(err)
			}
			msg.Ack(false)
			fmt.Printf("Worker %d has processed order %s\n", i, inputDTO.ID)
		}
	}
}
