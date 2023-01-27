package errors

import "errors"

var (
	InsufficientFunds = func() error {
		err := errors.New("Insufficient Funds")
		return err
	}
)
