package core

import "github.com/google/uuid"

type Wallet struct {
	ID      uuid.UUID `db:"wallet_id"`
	Balance int       `db:"balance"`
}
