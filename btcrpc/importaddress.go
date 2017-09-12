package btcrpc

import "errors"

// ImportAddress start watching for the addr
func (rpc *BTCRpc) ImportAddress(addr string, rescan bool) error {
	var params = make([]interface{}, 3)
	// The hex-encoded script (or address)
	params[0] = addr
	// (default="") An optional label
	params[1] = addr
	// (default=true) Rescan the wallet for transactions
	params[2] = rescan
	resp, err := rpc.call("importaddress", params)
	if err != nil {
		return err
	}
	if resp.Error.Message != "" {
		return errors.New(resp.Error.Message)
	}
	return nil
}
