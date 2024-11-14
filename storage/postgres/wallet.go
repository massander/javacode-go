package postgres

import (
	"context"
	"fmt"
	"wallet-api/storage"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var _ storage.WalletStorage = (*WalletStorage)(nil)

type WalletStorage struct {
	pool *pgxpool.Pool
}

func (s *WalletStorage) Deposit(ctx context.Context, walletID uuid.UUID, amount int) (int, error) {
	const op = "storage.postgres.wallet.Deposit"

	currentBalance, err := s.Balance(ctx, walletID)
	if err != nil {
		return -1, fmt.Errorf("%s: %w", op, err)
	}

	newBalance := currentBalance + amount

	err = s.updateBalance(ctx, walletID, newBalance)
	if err != nil {
		return -1, fmt.Errorf("%s: %w", op, err)
	}

	return newBalance, nil
}

func (s *WalletStorage) Whithdraw(ctx context.Context, walletID uuid.UUID, amount int) (int, error) {
	const op = "storage.postgres.wallet.Whithdraw"

	currentBalance, err := s.Balance(ctx, walletID)
	if err != nil {
		return -1, fmt.Errorf("%s: %w", op, err)
	}

	newBalance := currentBalance
	if amount > currentBalance {
		return -1, fmt.Errorf("%s: %w", op, storage.ErrorInsufficientBalance)
	}

	newBalance -= amount

	err = s.updateBalance(ctx, walletID, newBalance)
	if err != nil {
		return -1, fmt.Errorf("%s: %w", op, err)
	}

	return newBalance, nil
}

func (s *WalletStorage) updateBalance(ctx context.Context, walletID uuid.UUID, balance int) error {
	const op = "storage.postgres.wallet.updateBalance"

	_, err := s.pool.Exec(ctx, "UPDATE wallets SET balance = $1 WHERE wallet_id = $2", balance, walletID)
	if err != nil {
		return fmt.Errorf("%s: failed to update balance: %w", op, err)
	}
	return nil
}

func (s *WalletStorage) Balance(ctx context.Context, walletID uuid.UUID) (int, error) {
	const op = "storage.postgres.wallet.Balance"

	var balance int
	err := s.pool.QueryRow(ctx, "SELECT balance FROM wallets WHERE wallet_id = $1", walletID).Scan(&balance)
	if err != nil {

		if err == pgx.ErrNoRows {
			return -1, fmt.Errorf("%s: %w", op, storage.ErrorWalletNotFound)
		}

		return -1, fmt.Errorf("%s: execute statement: %w", op, err)
	}

	return balance, nil
}
