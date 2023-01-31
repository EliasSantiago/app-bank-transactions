package domain

type TransactionQueueService interface {
	Consumer()
}

type TransactionQueueUseCase interface {
	Consumer()
}

type TransactionQueueRepository interface {
}
