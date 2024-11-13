package core

import (
	"encoding/json"
	"fmt"
)

type WalletOperation struct {
	name string
}

var (
	WalletDeposit  = WalletOperation{name: "DEPOSIT"}
	WalletWithdraw = WalletOperation{name: "WITHDRAW"}
)

var walletOperations = map[string]WalletOperation{
	WalletDeposit.name:  WalletDeposit,
	WalletWithdraw.name: WalletWithdraw,
}

func parseWalletOperation(value string) (WalletOperation, error) {
	const op = "core.ParseOperation"

	operation, exists := walletOperations[value]
	if !exists {
		return WalletOperation{}, fmt.Errorf("invalid operation type %q", value)
	}

	return operation, nil
}

func (wo WalletOperation) Name() string {
	return wo.name
}

func (wo WalletOperation) String() string {
	return wo.name
}

func (wo WalletOperation) MarshalJSON() ([]byte, error) {
	return []byte(wo.name), nil
}

func (wo *WalletOperation) UnmarshalJSON(data []byte) error {
	var tmp string
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	operation, err := parseWalletOperation(string(tmp))
	if err != nil {
		return err
	}

	wo.name = operation.name
	return nil
}
