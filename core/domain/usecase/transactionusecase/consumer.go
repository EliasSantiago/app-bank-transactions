package transactionusecase

import (
	"encoding/json"
	"fmt"
	"strconv"

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
			status := false
			if inputDTO.Status == "Pendente" {
				status = false
			} else {
				status = true
			}
			strconv.ParseBool("true")
			if err != nil {
				panic(err)
			}
			store := &dto.CreateTransactionStore{
				ID:        inputDTO.ID,
				From:      inputDTO.From,
				To:        inputDTO.To,
				Value:     inputDTO.Value,
				Status:    status,
				CreatedAt: inputDTO.CreatedAt,
			}
			err = usecase.repository.Transfer(store)
			if err != nil {
				panic(err)
			}
			msg.Ack(false)
			fmt.Printf("Worker %d has processed order %s\n", i, store.ID)
		}
	}
}

// func worker(delivereMessage <-chan amqp.Delivery, dto, workerID int) {
// 	for msg := range delivereMessage {

// 		err := json.Unmarshal(msg.Body, inputDTO)
// 		if err != nil {
// 			panic(err)
// 		}
// 		outputDTO, err := usercase.repository.Transfer()
// 		if err != nil {
// 			panic(err)
// 		}
// 		msg.Ack(false)
// 		fmt.Printf("Worker %d has processed order %s\n", workerID, outputDTO.ID)
// 		time.Sleep(1 * time.Second)
// 	}
// }
