package rpcservice

import (
	"context"
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
		Client        *gui.Client
	}

	CreateTransactionHandler struct {
		BackendServer string
		Client        *gui.Client
	}

	InjectTransactionRequest struct {
		Rawtx string `json:"rawtx"`
	}

	InjectResponse struct {
		Txid string `json:"txid"`
	}

	InjectTransactionHandler struct {
		BackendServer string
		Client        *gui.Client
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

	output, err := h.Client.Transaction(req.Txid)
	if err != nil {
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

	output, err := h.Client.CreateTransaction(req)
	if err != nil {
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

	output, err := h.Client.InjectTransaction(req.Rawtx)
	if err != nil {
		return nil, ErrCustomise(err)
	}
	return output, nil
}
