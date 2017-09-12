package btcrpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// BTCRpc describes RPC server credentials
type BTCRpc struct {
	Addr string

	user     string
	password string

	debug bool
}

// New creates a new RPC server
func New(addr string, user string, password string, debug bool) *BTCRpc {
	return &BTCRpc{
		Addr:     addr,
		user:     user,
		password: password,
		debug:    debug,
	}
}

// call internal method, real RPC call
func (rpc *BTCRpc) call(method string, args []interface{}) (*rpcResponse, error) {
	data := rpcRequest{"1.0", "btxer", method, args}
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	if rpc.debug {
		fmt.Println(string(dataBytes))
	}
	body := bytes.NewReader(dataBytes)

	rpcAddr := rpc.Addr
	if !strings.Contains(rpcAddr, "http://") {
		rpcAddr = "http://" + rpcAddr
	}
	req, err := http.NewRequest("POST", rpcAddr, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "text/plain;")
	req.SetBasicAuth(rpc.user, rpc.password)

	httpResponse, err := http.DefaultClient.Do(req)
	if err != nil {
		// re-try
		httpResponse, err = http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}
	}
	defer httpResponse.Body.Close()
	respBytes, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, err
	}

	var resp rpcResponse
	err = json.Unmarshal(respBytes, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// IsAddrWatched returns true if btcAddr is Watched
func (rpc *BTCRpc) IsAddrWatched(addr string) (bool, error) {
	accounts, err := rpc.ListAccounts()
	if err != nil {
		return false, err
	}
	for _, acc := range *accounts {
		if acc.Name == addr {
			return true, nil
		}
	}
	return false, nil
}
