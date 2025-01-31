package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"wallet-api/storage"
)

type Storage struct {
	pool *pgxpool.Pool

	Wallet storage.WalletStorage
}

func New(URL string) (*Storage, error) {
	const op = "storage.postgres.New"

	ctx := context.TODO()

	pool, err := pgxpool.New(ctx, URL)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	pingContext, cancelPing := context.WithTimeout(ctx, time.Second*2)
	defer cancelPing()

	if err := pool.Ping(pingContext); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	storage := &Storage{
		pool: pool,
		Wallet: &WalletStorage{
			pool: pool,
		},
	}

	return storage, nil
}

func (s *Storage) Close() {
	s.pool.Close()
}
