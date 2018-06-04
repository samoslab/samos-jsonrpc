package rpcservice

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/intel-go/fastjson"
	"github.com/osamingo/jsonrpc"
	"github.com/samoslab/samos/src/gui" //http,json helpers
	"github.com/skycoin/skycoin/src/visor"
)

type (
	TransactionRequest struct {
		Txid string `json:"txid"`
	}
	TransactionHandler struct {
		BackendServer string
	}

	CreateTransactionHandler struct {
		BackendServer string
	}

	InjectTransactionRequest struct {
		Rawtx string `json:"rawtx"`
	}

	InjectResponse struct {
		Txid string `json:"txid"`
	}
	InjectTransactionHandler struct {
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

func (h CreateTransactionHandler) Name() string {
	return "transactionCreate"
}

func (h CreateTransactionHandler) Params() interface{} {
	return gui.CreateTransactionRequest{}
}

func (h CreateTransactionHandler) Result() interface{} {
	return gui.CreateTransactionResponse{}
}

func (h CreateTransactionHandler) ServeJSONRPC(c context.Context, params *fastjson.RawMessage) (interface{}, *jsonrpc.Error) {
	var req gui.CreateTransactionRequest
	if err := jsonrpc.Unmarshal(params, &req); err != nil {
		return nil, ErrCustomise(err)
	}
	fmt.Printf("req %+v\n", req)

	if req.Wallet.ID == "" || len(req.To) == 0 {
		return nil, jsonrpc.ErrInvalidParams()
	}

	url := fmt.Sprintf("%s/wallet/transaction", h.BackendServer)
	fmt.Printf("url %s\n", url)
	reqBody, _ := params.MarshalJSON()

	byteBody, err := SendJsonRequest("POST", url, reqBody)
	if err != nil {
		return nil, ErrCustomise(err)
	}
	output := gui.CreateTransactionResponse{}
	if err := json.Unmarshal(byteBody, &output); err != nil {
		return nil, ErrCustomise(err)
	}
	return output, nil
}

func (h InjectTransactionHandler) Name() string {
	return "transactionInject"
}

func (h InjectTransactionHandler) Params() interface{} {
	return InjectTransactionRequest{}
}

func (h InjectTransactionHandler) Result() interface{} {
	return nil
}

func (h InjectTransactionHandler) ServeJSONRPC(c context.Context, params *fastjson.RawMessage) (interface{}, *jsonrpc.Error) {
	var req InjectTransactionRequest
	if err := jsonrpc.Unmarshal(params, &req); err != nil {
		return nil, ErrCustomise(err)
	}
	fmt.Printf("req %+v\n", req)

	if req.Rawtx == "" {
		return nil, jsonrpc.ErrInvalidParams()
	}

	url := fmt.Sprintf("%s/injectTransaction", h.BackendServer)
	fmt.Printf("url %s\n", url)
	reqBody, _ := params.MarshalJSON()

	byteBody, err := SendJsonRequest("POST", url, reqBody)
	if err != nil {
		return nil, ErrCustomise(err)
	}
	output := InjectResponse{
		Txid: string(byteBody),
	}
	return output, nil
}
