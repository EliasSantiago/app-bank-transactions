package errors

import "errors"

var (
	InsufficientFunds = func() error {
		err := errors.New("Insufficient Funds")
		return err
	}
	GetBalance = func() error {
		err := errors.New("Error Get Balance")
		return err
	}
	CheckDataExists = func() error {
		err := errors.New("Failed to check if data exists")
		return err
	}
	WalletNotExist = func() error {
		err := errors.New("Wallet does not exist")
		return err
	}
	UpdateBalance = func() error {
		err := errors.New("Failed to update balance")
		return err
	}
	InvalidParameters = func() error {
		err := errors.New("Invalid parameters")
		return err
	}
	Unauthorized = func() error {
		err := errors.New("Not authorized by external service")
		return err
	}
)
