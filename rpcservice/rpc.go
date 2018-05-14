package rpcservice

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/intel-go/fastjson"
	"github.com/osamingo/jsonrpc"
	"github.com/skycoin/skycoin/src/visor"
	"github.com/skycoin/skycoin/src/wallet"
	"github.com/spaco/spo/src/cipher"
	"github.com/spaco/spo/src/gui"
)

type (
	CoinServer struct {
		Server string
	}
	StatusHandler  struct{}
	VersionHandler struct{}
	OutputsHandler struct{}

	OutputsRequest struct {
		Addrs  string `json:"addrs"`
		Hashes string `json:"hashes"`
	}

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

	AddressNewRequest struct {
		ID       string `json:"ID"`
		Num      int    `json:"num"`
		Password string `json:"password"`
	}

	WalletSpentRequest struct {
		ID       string `json:"ID"`
		Dst      string `json:"dst"`
		Coins    uint64 `json:"coins"`
		Password string `json:"password"`
	}

	EncryptWalletRequest struct {
		ID       string `json:"ID"`
		Password string `json:"password"`
	}

	BlockRequest struct {
		Seq  uint64 `json:"seq"`
		Hash string `json:"hash"`
	}
	BlockLastNRequest struct {
		Num uint64 `json:"num"`
	}

	BlockRangeRequest struct {
		Start uint64 `json:"start"`
		End   uint64 `json:"end"`
	}

	TransactionRequest struct {
		Txid string `json:"txid"`
	}

	WalletBalanceHandler struct{}
	BalanceHandler       struct{}

	WalletHandler       struct{}
	WalletCreateHandler struct{}
	WalletSpentHandler  struct{}

	AddressNewHandler struct{}

	EncryptWalletHandler struct{}
	DecryptWalletHandler struct{}

	BlockHandler      struct{}
	BlockRangeHandler struct{}
	BlockLastNHandler struct{}

	TransactionHandler struct{}
)

func (h AddressNewHandler) ServeJSONRPC(c context.Context, params *fastjson.RawMessage) (interface{}, *jsonrpc.Error) {
	var req AddressNewRequest
	if err := jsonrpc.Unmarshal(params, &req); err != nil {
		return nil, ErrCustomise(err)
	}
	if req.ID == "" || req.Password == "" {
		return nil, jsonrpc.ErrInvalidParams()
	}

	url := "http://127.0.0.1:6420/wallet/newAddress"
	fmt.Printf("url %s\n", url)
	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, ErrCustomise(err)
	}
	byteBody, err := SendRequest("POST", url, reqBody)
	if err != nil {
		return nil, ErrCustomise(err)
	}
	output := []cipher.Address{}
	if err := json.Unmarshal(byteBody, &output); err != nil {
		return nil, ErrCustomise(err)
	}
	return output, nil
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

func (h OutputsHandler) ServeJSONRPC(c context.Context, params *fastjson.RawMessage) (interface{}, *jsonrpc.Error) {
	req := &OutputsRequest{}
	if err := jsonrpc.Unmarshal(params, req); err != nil {
		return nil, ErrCustomise(err)
	}

	if req.Addrs != "" && req.Hashes != "" {
		return nil, jsonrpc.ErrInvalidParams()
	}

	url := "http://127.0.0.1:6420/outputs"
	if req.Addrs != "" {
		url = fmt.Sprintf("%s?addrs=%s", url, req.Addrs)
	} else if req.Hashes != "" {
		url = fmt.Sprintf("%s?hashes=%s", url, req.Hashes)
	}
	byteBody, err := SendRequest("GET", url, nil)
	if err != nil {
		return nil, ErrCustomise(err)
	}
	output := visor.ReadableOutputSet{}
	if err := json.Unmarshal(byteBody, &output); err != nil {
		return nil, ErrCustomise(err)
	}
	return output, nil

}

func (h VersionHandler) ServeJSONRPC(c context.Context, params *fastjson.RawMessage) (interface{}, *jsonrpc.Error) {
	url := "http://127.0.0.1:6420/version"
	byteBody, err := SendRequest("GET", url, nil)
	if err != nil {
		return nil, ErrCustomise(err)
	}

	bi := visor.BuildInfo{}
	if err := json.Unmarshal(byteBody, &bi); err != nil {
		return nil, ErrCustomise(err)
	}

	return bi, nil
}

func SendRequest(method, url string, reqBody []byte) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewReader(reqBody))
	if err != nil {
		return nil, err
	}

	rsp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()
	byteBody, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	return byteBody, nil
}

// ErrCustomise returns invalid error by error.
func ErrCustomise(err error) *jsonrpc.Error {
	return &jsonrpc.Error{
		Code:    jsonrpc.ErrorCodeInternal,
		Message: err.Error(),
	}
}

func RegisterMethod() *jsonrpc.MethodRepository {

	mr := jsonrpc.NewMethodRepository()

	if err := mr.RegisterMethod("version", VersionHandler{}, struct{}{}, visor.BuildInfo{}); err != nil {
		log.Fatalln(err)
	}
	if err := mr.RegisterMethod("outputs", OutputsHandler{}, OutputsRequest{}, visor.ReadableOutputSet{}); err != nil {
		log.Fatalln(err)
	}
	if err := mr.RegisterMethod("walletBalance", WalletBalanceHandler{}, OnlyIDRequest{}, wallet.BalancePair{}); err != nil {
		log.Fatalln(err)
	}

	if err := mr.RegisterMethod("balance", BalanceHandler{}, BalanceRequest{}, wallet.BalancePair{}); err != nil {
		log.Fatalln(err)
	}

	if err := mr.RegisterMethod("wallet", WalletHandler{}, OnlyIDRequest{}, wallet.Wallet{}); err != nil {
		log.Fatalln(err)
	}

	if err := mr.RegisterMethod("walletCreate", WalletCreateHandler{}, WalletCreateRequest{}, wallet.Wallet{}); err != nil {
		log.Fatalln(err)
	}

	if err := mr.RegisterMethod("walletSpent", WalletSpentHandler{}, WalletSpentRequest{}, gui.SpendResult{}); err != nil {
		log.Fatalln(err)
	}

	if err := mr.RegisterMethod("addressNew", AddressNewHandler{}, AddressNewRequest{}, []cipher.Address{}); err != nil {
		log.Fatalln(err)
	}

	if err := mr.RegisterMethod("encryptWallet", EncryptWalletHandler{}, EncryptWalletRequest{}, wallet.Wallet{}); err != nil {
		log.Fatalln(err)
	}
	if err := mr.RegisterMethod("decryptWallet", DecryptWalletHandler{}, EncryptWalletRequest{}, wallet.Wallet{}); err != nil {
		log.Fatalln(err)
	}
	if err := mr.RegisterMethod("block", BlockHandler{}, BlockRequest{}, wallet.Wallet{}); err != nil {
		log.Fatalln(err)
	}

	return mr

}
