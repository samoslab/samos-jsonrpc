package rpcservice

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/intel-go/fastjson"
	"github.com/osamingo/jsonrpc"
	"github.com/skycoin/skycoin/src/gui"
	"github.com/skycoin/skycoin/src/wallet"
)

type (
	OnlyIDRequest struct {
		ID string `json:"id"`
	}
	BalanceRequest struct {
		Addrs string `json:"addrs"`
	}

	WalletCreateRequest struct {
		Seed     string `json:"seed"`
		Lable    string `json:"label"`
		Scan     int    `json:"scan"`
		Password string `json:"password"`
	}

	WalletSpentRequest struct {
		ID       string `json:"ID"`
		Dst      string `json:"dst"`
		Coins    uint64 `json:"coins"`
		Password string `json:"password"`
	}

	WalletEncryptRequest struct {
		ID       string `json:"ID"`
		Password string `json:"password"`
	}
	BalanceHandler       struct{}
	WalletBalanceHandler struct{}

	WalletHandler       struct{}
	WalletCreateHandler struct{}
	WalletSpentHandler  struct{}

	WalletEncryptHandler struct{}
	WalletDecryptHandler struct{}
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

func (h WalletSpentHandler) Name() string {
	return "walletSpend"
}

func (h WalletSpentHandler) Params() interface{} {
	return WalletSpentRequest{}
}

func (h WalletSpentHandler) Result() interface{} {
	return gui.SpendResult{}
}

func (h WalletSpentHandler) ServeJSONRPC(c context.Context, params *fastjson.RawMessage) (interface{}, *jsonrpc.Error) {
	var req WalletSpentRequest
	if err := jsonrpc.Unmarshal(params, &req); err != nil {
		return nil, ErrCustomise(err)
	}
	if req.ID == "" || req.Dst == "" {
		return nil, jsonrpc.ErrInvalidParams()
	}

	url := "http://127.0.0.1:6420/wallet/spend"
	fmt.Printf("url %s\n", url)
	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, ErrCustomise(err)
	}
	byteBody, err := SendRequest("POST", url, reqBody)
	if err != nil {
		return nil, ErrCustomise(err)
	}
	output := &gui.SpendResult{}
	if err := json.Unmarshal(byteBody, &output); err != nil {
		return nil, ErrCustomise(err)
	}
	return output, nil
}

func (h WalletCreateHandler) Name() string {
	return "walletCreate"
}

func (h WalletCreateHandler) Params() interface{} {
	return WalletCreateRequest{}
}

func (h WalletCreateHandler) Result() interface{} {
	return wallet.Wallet{}
}
func (h WalletCreateHandler) ServeJSONRPC(c context.Context, params *fastjson.RawMessage) (interface{}, *jsonrpc.Error) {
	var req WalletCreateRequest
	if err := jsonrpc.Unmarshal(params, &req); err != nil {
		return nil, ErrCustomise(err)
	}
	if req.Seed == "" || req.Lable == "" {
		return nil, jsonrpc.ErrInvalidParams()
	}

	url := "http://127.0.0.1:6420/wallet/create"
	fmt.Printf("url %s\n", url)
	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, ErrCustomise(err)
	}
	byteBody, err := SendRequest("POST", url, reqBody)
	if err != nil {
		return nil, ErrCustomise(err)
	}
	output := &wallet.Wallet{}
	if err := json.Unmarshal(byteBody, &output); err != nil {
		return nil, ErrCustomise(err)
	}
	return output, nil
}

func (h WalletHandler) Name() string {
	return "wallet"
}

func (h WalletHandler) Params() interface{} {
	return OnlyIDRequest{}
}

func (h WalletHandler) Result() interface{} {
	return wallet.Wallet{}
}
func (h WalletHandler) ServeJSONRPC(c context.Context, params *fastjson.RawMessage) (interface{}, *jsonrpc.Error) {
	var req OnlyIDRequest
	if err := jsonrpc.Unmarshal(params, &req); err != nil {
		return nil, ErrCustomise(err)
	}
	if req.ID == "" {
		return nil, jsonrpc.ErrInvalidParams()
	}

	url := fmt.Sprintf("http://127.0.0.1:6420/wallet?id=%s", req.ID)
	fmt.Printf("url %s\n", url)
	byteBody, err := SendRequest("GET", url, nil)
	if err != nil {
		return nil, ErrCustomise(err)
	}
	output := &wallet.Wallet{}
	if err := json.Unmarshal(byteBody, &output); err != nil {
		return nil, ErrCustomise(err)
	}
	return output, nil
}

func (h WalletEncryptHandler) Name() string {
	return "walletEncrypt"
}

func (h WalletEncryptHandler) Params() interface{} {
	return WalletEncryptRequest{}
}

func (h WalletEncryptHandler) Result() interface{} {
	return wallet.Wallet{}
}
func (h WalletEncryptHandler) ServeJSONRPC(c context.Context, params *fastjson.RawMessage) (interface{}, *jsonrpc.Error) {
	var req WalletEncryptRequest
	if err := jsonrpc.Unmarshal(params, &req); err != nil {
		return nil, ErrCustomise(err)
	}
	if req.ID == "" {
		return nil, jsonrpc.ErrInvalidParams()
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, ErrCustomise(err)
	}
	url := "http://127.0.0.1:6420/wallet/encrypt"
	fmt.Printf("url %s\n", url)
	byteBody, err := SendRequest("POST", url, reqBody)
	if err != nil {
		return nil, ErrCustomise(err)
	}
	output := &wallet.Wallet{}
	if err := json.Unmarshal(byteBody, &output); err != nil {
		return nil, ErrCustomise(err)
	}
	return output, nil
}

func (h WalletDecryptHandler) Name() string {
	return "walletDecrypt"
}

func (h WalletDecryptHandler) Params() interface{} {
	return WalletEncryptRequest{}
}

func (h WalletDecryptHandler) Result() interface{} {
	return wallet.Wallet{}
}
func (h WalletDecryptHandler) ServeJSONRPC(c context.Context, params *fastjson.RawMessage) (interface{}, *jsonrpc.Error) {
	var req WalletEncryptRequest
	if err := jsonrpc.Unmarshal(params, &req); err != nil {
		return nil, ErrCustomise(err)
	}
	if req.ID == "" {
		return nil, jsonrpc.ErrInvalidParams()
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, ErrCustomise(err)
	}
	url := "http://127.0.0.1:6420/wallet/decrypt"
	byteBody, err := SendRequest("POST", url, reqBody)
	if err != nil {
		return nil, ErrCustomise(err)
	}
	output := &wallet.Wallet{}
	if err := json.Unmarshal(byteBody, &output); err != nil {
		return nil, ErrCustomise(err)
	}
	return output, nil
}
