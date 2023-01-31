package transactionservice

func (service service) Consumer() {
	service.usecase.Consumer()
}
