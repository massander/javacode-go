package main

import (
	"fmt"
	"math/rand"

	"github.com/google/uuid"
)

func seed() {
	// wallets := make([]core.Wallet, 100)
	// rand.Seed(time.Now().UnixNano())

	for range 100 {
		// wallets = append(wallets, core.Wallet{
		// 	ID:      uuid.New(),
		// 	Balance: rand.Intn(100001),
		// })

		fmt.Printf("('%s', '%d'),\n", uuid.New().String(), rand.Intn(100))
	}

}

func main() {
	seed()
}
