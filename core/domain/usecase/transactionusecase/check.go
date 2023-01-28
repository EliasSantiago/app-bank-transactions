package transactionusecase

func (usecase usecase) Check(id string) (bool, error) {
	exist, err := usecase.repository.Check(id)
	if err != nil {
		return false, err
	}
	return exist, nil
}
