package btcrpc

import "errors"
import "encoding/json"

// GetBalance return balance for the addr
func (rpc *BTCRpc) GetBalance(addr string) (float32, error) {
	var params = make([]interface{}, 3)
	// The account string
	params[0] = addr
	// (default=1) Only include transactions confirmed at least this many times
	params[1] = 1
	// (default=false) Also include balance in watch-only addresses (see 'importaddress')
	params[2] = true
	resp, err := rpc.call("getbalance", params)
	if err != nil {
		return 0, err
	}
	if resp.Error.Message != "" {
		return 0, errors.New(resp.Error.Message)
	}
	var balance float32
	if err = json.Unmarshal(resp.Result, &balance); err != nil {
		return 0, err
	}
	return balance, nil
}
