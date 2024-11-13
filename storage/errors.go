package storage

import "errors"

var (
	ErrorInsufficientBalance = errors.New("insufficient balance")
	ErrorWalletNotFound      = errors.New("wallet not found")
)
