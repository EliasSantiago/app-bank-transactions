package transactionservice

import (
	"encoding/json"
	"net/http"

	"github.com/EliasSantiago/app-bank-transactions/adapter/http/errors"
	"github.com/EliasSantiago/app-bank-transactions/core/dto"
)

// @Summary Create new transaction
// @Description Create new transaction
// @Tags transaction
// @Accept  json
// @Produce  json
// @Param transaction body dto.CreateTransactionRequest true "transaction"
// @Success 200 {object} domain.Transaction
// @Router /transaction [post]
func (service service) Create(response http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	transactionRequest, err := dto.FromJSONCreateTransactionRequest(request.Body)
	if err != nil {
		errors.InvalidParameters()
		return
	}
	transaction, err := service.usecase.Create(transactionRequest)
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(response).Encode(transaction)
}
