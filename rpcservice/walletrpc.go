package rpcservice

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/intel-go/fastjson"
	"github.com/osamingo/jsonrpc"
	"github.com/skycoin/skycoin/src/wallet"
)

type (
	OnlyIDRequest struct {
		ID string `json:"id"`
	}
	BalanceRequest struct {
		Addrs string `json:"addrs"`
	}
	BalanceHandler       struct{}
	WalletBalanceHandler struct{}
)

func (h BalanceHandler) Name() string {
	return "balance"
}

func (h BalanceHandler) Params() interface{} {
	return BalanceRequest{}
}

func (h BalanceHandler) Result() interface{} {
	return wallet.BalancePair{}
}

func (h BalanceHandler) ServeJSONRPC(c context.Context, params *fastjson.RawMessage) (interface{}, *jsonrpc.Error) {
	var req BalanceRequest
	if err := jsonrpc.Unmarshal(params, &req); err != nil {
		return nil, ErrCustomise(err)
	}
	if req.Addrs == "" {
		return nil, jsonrpc.ErrInvalidParams()
	}

	url := fmt.Sprintf("http://127.0.0.1:6420/balance?addrs=%s", req.Addrs)
	fmt.Printf("url %s\n", url)
	byteBody, err := SendRequest("GET", url, nil)
	if err != nil {
		return nil, ErrCustomise(err)
	}
	output := wallet.BalancePair{}
	if err := json.Unmarshal(byteBody, &output); err != nil {
		return nil, ErrCustomise(err)
	}
	return output, nil
}

func (h WalletBalanceHandler) Name() string {
	return "walletBalance"
}

func (h WalletBalanceHandler) Params() interface{} {
	return OnlyIDRequest{}
}

func (h WalletBalanceHandler) Result() interface{} {
	return wallet.BalancePair{}
}

func (h WalletBalanceHandler) ServeJSONRPC(c context.Context, params *fastjson.RawMessage) (interface{}, *jsonrpc.Error) {
	var req OnlyIDRequest
	if err := jsonrpc.Unmarshal(params, &req); err != nil {
		return nil, ErrCustomise(err)
	}
	fmt.Printf("id %s\n", req.ID)
	if req.ID == "" {
		return nil, jsonrpc.ErrInvalidParams()
	}

	url := fmt.Sprintf("http://127.0.0.1:6420/wallet/balance?id=%s", req.ID)
	fmt.Printf("url %s\n", url)
	byteBody, err := SendRequest("GET", url, nil)
	if err != nil {
		return nil, ErrCustomise(err)
	}
	output := wallet.BalancePair{}
	if err := json.Unmarshal(byteBody, &output); err != nil {
		return nil, ErrCustomise(err)
	}
	return output, nil
}
