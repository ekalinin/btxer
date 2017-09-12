package btcrpc

import (
	"encoding/json"
	"errors"
)

// BTCTx describes a transaction
type BTCTx struct {
	Category  string
	Amount    float32
	Label     string
	Vout      int
	Fee       float32
	Txid      string
	Blocktime int64
}

// ListTransactions return list of transactions for addr
func (rpc *BTCRpc) ListTransactions(addr string, count int, skip int) ([]BTCTx, error) {

	var params = make([]interface{}, 4)
	// The account name
	params[0] = addr
	// (default=10) The number of transactions to return
	params[1] = count
	// (default=0) The number of transactions to skip
	params[2] = skip
	// (default=false) Include transactions to watch-only addresses
	params[3] = true
	resp, err := rpc.call("listtransactions", params)
	if err != nil {
		return nil, err
	}
	if resp.Error.Message != "" {
		return nil, errors.New(resp.Error.Message)
	}
	var txs = []BTCTx{}
	if err = json.Unmarshal(resp.Result, &txs); err != nil {
		return nil, err
	}
	return txs, nil
}
