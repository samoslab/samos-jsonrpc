package rpcservice

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/intel-go/fastjson"
	"github.com/osamingo/jsonrpc"
	"github.com/skycoin/skycoin/src/visor"
)

type (
	TransactionRequest struct {
		Txid string `json:"txid"`
	}
	TransactionHandler struct {
		BackendServer string
	}
)

func (h TransactionHandler) Name() string {
	return "transaction"
}

func (h TransactionHandler) Params() interface{} {
	return TransactionRequest{}
}

func (h TransactionHandler) Result() interface{} {
	return visor.TransactionResult{}
}

func (h TransactionHandler) ServeJSONRPC(c context.Context, params *fastjson.RawMessage) (interface{}, *jsonrpc.Error) {
	var req TransactionRequest
	if err := jsonrpc.Unmarshal(params, &req); err != nil {
		return nil, ErrCustomise(err)
	}

	if req.Txid == "" {
		return nil, jsonrpc.ErrInvalidParams()
	}

	url := fmt.Sprintf("%s/transaction?txid=%s", h.BackendServer, req.Txid)
	fmt.Printf("url %s\n", url)
	byteBody, err := SendRequest("GET", url, nil)
	if err != nil {
		return nil, ErrCustomise(err)
	}
	output := visor.TransactionResult{}
	if err := json.Unmarshal(byteBody, &output); err != nil {
		return nil, ErrCustomise(err)
	}
	return output, nil
}
