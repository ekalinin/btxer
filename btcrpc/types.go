package btcrpc

import "encoding/json"

type rpcError struct {
	Code    int `json:"code,omitempty"`
	Message string       `json:"message,omitempty"`
}

type rpcResponse struct {
	Result json.RawMessage `json:"result"`
	Error  rpcError        `json:"error"`
}

type rpcRequest struct {
	Jsonrpc string        `json:"jsonrpc"`
	ID      string        `json:"id"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}
