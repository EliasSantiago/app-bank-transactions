package dto_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/EliasSantiago/app-bank-transactions/core/dto"
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/require"
)

func TestFromJSONCreateWalletRequest(t *testing.T) {
	fakeItem := dto.CreateWalletRequest{}
	faker.FakeData(&fakeItem)
	json, err := json.Marshal(fakeItem)
	require.Nil(t, err)
	itemRequest, err := dto.FromJSONCreateWalletRequest(strings.NewReader(string(json)))
	require.Nil(t, err)
	require.Equal(t, itemRequest.UserID, fakeItem.UserID)
	require.Equal(t, itemRequest.Balance, fakeItem.Balance)
}

func TestFromJSONCreateWalletRequest_JSONDecodeError(t *testing.T) {
	itemRequest, err := dto.FromJSONCreateWalletRequest(strings.NewReader("{"))
	require.NotNil(t, err)
	require.Nil(t, itemRequest)
}
