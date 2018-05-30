package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/intel-go/fastjson"
	"github.com/osamingo/jsonrpc"
	"github.com/samoslab/sky-fiber-jsonrpc/rpcservice"
)

type (
	HandleParamsResulter interface {
		jsonrpc.Handler
		Name() string
		Params() interface{}
		Result() interface{}
	}
)

type (
	EchoHandler struct{}
	EchoParams  struct {
		Name string `json:"name"`
	}
	EchoResult struct {
		Message string `json:"message"`
	}
)

func (h EchoHandler) Name() string {
	return "Echo"
}

func (h EchoHandler) Params() interface{} {
	return EchoParams{}
}

func (h EchoHandler) Result() interface{} {
	return EchoResult{}
}

func (h EchoHandler) ServeJSONRPC(c context.Context, params *fastjson.RawMessage) (interface{}, *jsonrpc.Error) {

	var p EchoParams
	if err := jsonrpc.Unmarshal(params, &p); err != nil {
		return nil, err
	}

	return EchoResult{
		Message: "Hello, " + p.Name,
	}, nil
}

func NewHandlers() []HandleParamsResulter {
	return []HandleParamsResulter{
		EchoHandler{},
		rpcservice.BalanceHandler{},
		rpcservice.WalletBalanceHandler{},
		rpcservice.WalletSpentHandler{},
		rpcservice.WalletCreateHandler{},
		rpcservice.WalletHandler{},
		rpcservice.WalletEncryptHandler{},
		rpcservice.WalletDecryptHandler{},

		rpcservice.VersionHandler{},
		rpcservice.AddressNewHandler{},
		rpcservice.OutputsHandler{},
		rpcservice.BlockHandler{},
		rpcservice.BlockRangeHandler{},
		rpcservice.BlockLastNHandler{},
		rpcservice.TransactionHandler{},
	}
}

func main() {

	mr := jsonrpc.NewMethodRepository()

	for _, h := range NewHandlers() {
		fmt.Printf("%s\n", h.Name())
		mr.RegisterMethod(h.Name(), h, h.Params(), h.Result())
	}

	http.Handle("/jrpc", mr)
	http.HandleFunc("/jrpc/debug", mr.ServeDebug)

	if err := http.ListenAndServe(":8080", http.DefaultServeMux); err != nil {
		log.Fatalln(err)
	}
}
