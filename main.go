package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/patrickmn/go-cache"

	api "wallet-api/api/v1"
	"wallet-api/storage/postgres"
)

func main() {
	db, err := postgres.New(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	cache := cache.New(5*time.Minute, 10*time.Minute)

	mux := http.NewServeMux()

	apiv1 := api.NewAPIv1Service(db, cache)
	apiv1.RegisterGateway(mux)

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
