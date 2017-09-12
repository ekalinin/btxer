package btcrpc

import (
	"encoding/json"
	"errors"
	"strconv"
)

// BTCAccount describes a BTC Account
type BTCAccount struct {
	Name    string
	Balance string
}

// BTCAccounts a list of BTCAccount
type BTCAccounts []BTCAccount

// ListAccounts return accounts
func (rpc *BTCRpc) ListAccounts() (*BTCAccounts, error) {
	var params = make([]interface{}, 2)
	// Only include transactions with at least this many confirmations
	params[0] = 1
	// Include balances in watch-only addresses (see 'importaddress')
	params[1] = true
	resp, err := rpc.call("listaccounts", params)
	if err != nil {
		return nil, err
	}
	if resp.Error.Message != "" {
		return nil, errors.New(resp.Error.Message)
	}

	var tmpAccounts map[string]float64
	if err = json.Unmarshal(resp.Result, &tmpAccounts); err != nil {
		return nil, err
	}

	accounts := make(BTCAccounts, len(tmpAccounts))
	for k, v := range tmpAccounts {
		accounts = append(accounts, BTCAccount{k, strconv.FormatFloat(v, 'f', -1, 64)})
	}
	return &accounts, nil
}
