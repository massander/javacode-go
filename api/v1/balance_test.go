package v1

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	gocache "github.com/patrickmn/go-cache"
	"github.com/stretchr/testify/require"

	"wallet-api/storage/postgres"
)

func Test_APIv1Service_updateBalance(t *testing.T) {
	cases := []struct {
		name      string
		walletID  string
		operation string
		amount    int
		balance   int
		errcode   string
		err       string
	}{
		{
			name:     "OK",
			walletID: "3624351d-ddc0-4e93-8eed-bd091bb4c7f1",
			amount:   50,
			balance:  150,
			errcode:  "",
			err:      "",
		},
	}

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	db, err := postgres.New(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	cache := gocache.New(5*time.Minute, 10*time.Minute)

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			input := fmt.Sprintf(`{"walletId": "%s", "amount": "%s", "operationType": "%s"}`, tc.url, tc.alias)

			mux := http.NewServeMux()

			apiv1 := NewAPIv1Service(db, cache)
			apiv1.RegisterGateway(mux)

			req, err := http.NewRequest(http.MethodPost, "/save", bytes.NewReader([]byte(input)))
			require.NoError(t, err)

			rr := httptest.NewRecorder()

		})

	}

}
