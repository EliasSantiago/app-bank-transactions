package walletservice

import (
	"encoding/json"
	"net/http"

	"github.com/EliasSantiago/app-bank-transactions/adapter/http/errors"
	"github.com/EliasSantiago/app-bank-transactions/core/dto"
)

// @Summary Create new wallet
// @Description Create new wallet
// @Tags wallet
// @Accept  json
// @Produce  json
// @Param user body dto.CreateWalletRequest true "wallet"
// @Success 200 {object} domain.Wallet
// @Router /wallet [post]
func (service service) Create(response http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	walletRequest, err := dto.FromJSONCreateWalletRequest(request.Body)
	if err != nil {
		errors.InvalidParameters()
		return
	}
	wallet, err := service.usecase.Create(walletRequest)
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(response).Encode(wallet)
}
