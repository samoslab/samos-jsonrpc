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

	WalletBalanceHandler struct{}
	BalanceHandler       struct{}

	WalletHandler struct{}
)

func (h WalletHandler) ServeJSONRPC(c context.Context, params *fastjson.RawMessage) (interface{}, *jsonrpc.Error) {
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

func RegisterMethod() *jsonrpc.MethodRepository {

	mr := jsonrpc.NewMethodRepository()

	if err := mr.RegisterMethod("version", VersionHandler{}, struct{}{}, visor.BuildInfo{}); err != nil {
		log.Fatalln(err)
	}
	if err := mr.RegisterMethod("outputs", OutputsHandler{}, OutputsRequest{}, visor.ReadableOutputSet{}); err != nil {
		log.Fatalln(err)
	}
	if err := mr.RegisterMethod("walletBalance", WalletBalanceHandler{}, new(string), wallet.BalancePair{}); err != nil {
		log.Fatalln(err)
	}

	if err := mr.RegisterMethod("balance", BalanceHandler{}, BalanceRequest{}, wallet.BalancePair{}); err != nil {
		log.Fatalln(err)
	}

	return mr

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
