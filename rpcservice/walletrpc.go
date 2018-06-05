package rpcservice

import (
	"context"
	"fmt"

	"github.com/intel-go/fastjson"
	"github.com/osamingo/jsonrpc"
	"github.com/samoslab/samos/src/gui"
	"github.com/skycoin/skycoin/src/wallet"
)

type (
	OnlyIDRequest struct {
		ID string `json:"id"`
	}
	BalanceRequest struct {
		Addrs []string `json:"addrs"`
	}

	WalletCreateRequest struct {
		Seed     string `json:"seed"`
		Label    string `json:"label"`
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
	BalanceHandler struct {
		BackendServer string
		Client        *gui.Client
	}
	WalletBalanceHandler struct {
		BackendServer string
		Client        *gui.Client
	}

	WalletHandler struct {
		BackendServer string
		Client        *gui.Client
	}
	WalletCreateHandler struct {
		BackendServer string
		Client        *gui.Client
	}
	WalletSpentHandler struct {
		BackendServer string
		Client        *gui.Client
	}

	WalletEncryptHandler struct {
		BackendServer string
		Client        *gui.Client
	}
	WalletDecryptHandler struct {
		BackendServer string
		Client        *gui.Client
	}
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
	if len(req.Addrs) == 0 {
		return nil, jsonrpc.ErrInvalidParams()
	}

	output, err := h.Client.Balance(req.Addrs)
	if err != nil {
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

	output, err := h.Client.WalletBalance(req.ID)
	if err != nil {
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

	output, err := h.Client.Spend(req.ID, req.Dst, req.Coins, req.Password)
	if err != nil {
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
	return gui.WalletResponse{}
}
func (h WalletCreateHandler) ServeJSONRPC(c context.Context, params *fastjson.RawMessage) (interface{}, *jsonrpc.Error) {
	var req WalletCreateRequest
	if err := jsonrpc.Unmarshal(params, &req); err != nil {
		return nil, ErrCustomise(err)
	}
	fmt.Printf("%v\n", req)
	if req.Seed == "" || req.Label == "" {
		return nil, jsonrpc.ErrInvalidParams()
	}

	output, err := h.Client.CreateEncryptedWallet(req.Seed, req.Label, req.Password, req.Scan)
	if err != nil {
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

	output, err := h.Client.Wallet(req.ID)
	if err != nil {
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
	output, err := h.Client.EncryptWallet(req.ID, req.Password)
	if err != nil {
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

	output, err := h.Client.DecryptWallet(req.ID, req.Password)
	if err != nil {
		return nil, ErrCustomise(err)
	}
	return output, nil
}
