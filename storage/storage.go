package storage

import (
	"context"
	"wallet-api/core"

	"github.com/google/uuid"
)

type BaseStorage[T any] interface {
	// Save(ctx context.Context, entity T) (T, error)
	// FindAll(ctx context.Context) ([]T, error)
	// FindById(ctx context.Context, id int) (T, error)
	// Delete(ctx context.Context, id int) error
}

type WalletStorage interface {
	BaseStorage[core.Wallet]
	Deposit(ctx context.Context, walletID uuid.UUID, amount int) (int, error)
	Whithdraw(ctx context.Context, walletID uuid.UUID, amount int) (int, error)
	Balance(ctx context.Context, walletID uuid.UUID) (int, error)
}
