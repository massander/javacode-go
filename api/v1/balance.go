package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"

	"wallet-api/core"
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

func (s *APIv1Service) updateBalance(w http.ResponseWriter, r *http.Request) {
	var payload updateBalanceRequest
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// TODO: get balance

	switch payload.OperationType {
	case core.WalletDeposit:
		// Do stuff
	case core.WalletWithdraw:
		// Do staff
	default:
		http.Error(w, "Invalid operation type", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(walletResponse{Balance: 0, WalletID: payload.WalletID})
	return

}

func (s *APIv1Service) getBalance(w http.ResponseWriter, r *http.Request) {
	walletID, err := uuid.Parse(r.PathValue("WALLET_UUID"))
	if err != nil {
		http.Error(w, "Invalid wallet ID", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(walletResponse{Balance: 0, WalletID: walletID})
	return
}
