package v1

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	gocache "github.com/patrickmn/go-cache"

	"wallet-api/core"
	"wallet-api/storage"
)

type updateBalanceRequest struct {
	WalletID      uuid.UUID            `json:"walletId"`
	OperationType core.WalletOperation `json:"operationType"`
	Amount        int                  `json:"amount"`
}

type walletResponse struct {
	WalletID uuid.UUID `json:"walletId"`
	Balance  int       `json:"balance"`
}

type errorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (s *APIv1Service) updateBalance(w http.ResponseWriter, r *http.Request) {
	var payload updateBalanceRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		sendResponse(w, r, http.StatusBadRequest,
			errorResponse{
				Code: "INVALID_PAYLOAD",
			},
		)
		return
	}

	ctx := r.Context()

	var balance int
	var err error

	switch payload.OperationType {
	case core.WalletDeposit:
		balance, err = s.storage.Wallet.Deposit(ctx, payload.WalletID, payload.Amount)
	case core.WalletWithdraw:
		balance, err = s.storage.Wallet.Whithdraw(ctx, payload.WalletID, payload.Amount)
	default:
		sendResponse(w, r, http.StatusBadRequest,
			errorResponse{
				Code:    "INVALID_OPERATION_TYPE",
				Message: "Operation type must be DEPOSIT or WHITHDRAW.",
			},
		)
		return
	}

	if err != nil {
		if errors.Is(err, storage.ErrorInsufficientBalance) {
			sendResponse(w, r, http.StatusConflict,
				errorResponse{
					Code:    "INSUFFICIENT_BALANCE",
					Message: "Please check your wallet and ensure sufficient funds before trying again.",
				},
			)
			return
		}

		if errors.Is(err, storage.ErrorWalletNotFound) {
			sendResponse(w, r, http.StatusNotFound,
				errorResponse{
					Code:    "INVALID_WALLET",
					Message: "Check that your requests are specifying a valid wallet ID.",
				},
			)
			return
		}

		sendResponse(w, r, http.StatusNotFound,
			errorResponse{
				Code:    "PROCESSING_ERROR",
				Message: "An error occurred while processing the operation.",
			},
		)
		return
	}

	s.cache.Set(payload.WalletID.String(), balance, gocache.DefaultExpiration)

	sendResponse(w, r, http.StatusOK,
		walletResponse{
			Balance:  balance,
			WalletID: payload.WalletID,
		},
	)

}

func (s *APIv1Service) getBalance(w http.ResponseWriter, r *http.Request) {
	walletID, err := uuid.Parse(r.PathValue("WALLET_UUID"))
	if err != nil {
		sendResponse(w, r, http.StatusNotFound,
			errorResponse{
				Code:    "INVALID_WALLET_PARAMETR",
				Message: "Check that your requests are specifying a valid wallet ID in the URL.",
			},
		)
		return
	}

	var balance int

	if x, found := s.cache.Get(walletID.String()); found {
		balance = x.(int)
	} else {
		balance, err = s.storage.Wallet.Balance(r.Context(), walletID)
		if err != nil {
			fmt.Println(err)
			if errors.Is(err, storage.ErrorWalletNotFound) {
				sendResponse(w, r, http.StatusNotFound,
					errorResponse{
						Code:    "INVALID_WALLET",
						Message: "Check that your requests are specifying a valid wallet ID.",
					},
				)
				return
			}
			sendResponse(w, r, http.StatusNotFound,
				errorResponse{
					Code:    "PROCESSING_ERROR",
					Message: "An error occurred while processing the operation.",
				},
			)
			return
		}

		s.cache.Set(walletID.String(), balance, gocache.DefaultExpiration)
	}

	sendResponse(w, r, http.StatusOK,
		walletResponse{
			Balance:  balance,
			WalletID: walletID,
		},
	)
	return
}
