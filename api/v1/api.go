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

func (s *APIv1Service) UpdateBalance(w http.ResponseWriter, r *http.Request) {

}

func (s *APIv1Service) GetBalance(w http.ResponseWriter, r *http.Request) {

}
