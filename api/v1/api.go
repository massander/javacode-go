package v1

import (
	"net/http"
	"wallet-api/storage/postgres"

	gocache "github.com/patrickmn/go-cache"
)

type APIv1Service struct {
	storage *postgres.Storage
	cache   *gocache.Cache
}

func NewAPIv1Service(storage *postgres.Storage, cache *gocache.Cache) *APIv1Service {
	api := &APIv1Service{
		storage: storage,
		cache:   cache,
	}

	return api
}

func (s *APIv1Service) RegisterGateway(mux *http.ServeMux) error {
	mux.HandleFunc("POST /api/v1/wallet", s.updateBalance)
	mux.HandleFunc("GET /api/v1/wallets/{WALLET_UUID}", s.getBalance)

	return nil
}
