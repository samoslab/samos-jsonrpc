package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/intel-go/fastjson"
	"github.com/osamingo/jsonrpc"
	"github.com/samoslab/skyjsonrpc/rpcservice"
)

type (
	HandleParamsResulter interface {
		jsonrpc.Handler
		Name() string
		Params() interface{}
		Result() interface{}
	}
	Servicer interface {
		MethodName(HandleParamsResulter) string
		Handlers() []HandleParamsResulter
	}
	UserService struct {
		SignUpHandler HandleParamsResulter
	}
	BalanceService struct {
		BalanceHandler       HandleParamsResulter
		WalletBalanceHandler HandleParamsResulter
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

func (us *UserService) MethodName(h HandleParamsResulter) string {
	return "UserService." + h.Name()
}

func (us *UserService) Handlers() []HandleParamsResulter {
	return []HandleParamsResulter{us.SignUpHandler}
}

func NewUserService() *UserService {
	return &UserService{
		SignUpHandler: EchoHandler{},
	}
}

func (us *BalanceService) MethodName(h HandleParamsResulter) string {
	return "BalanceService." + h.Name()
}

func (us *BalanceService) Handlers() []HandleParamsResulter {
	return []HandleParamsResulter{us.BalanceHandler, us.WalletBalanceHandler}
}

func NewBalanceService() *BalanceService {
	return &BalanceService{
		BalanceHandler:       rpcservice.BalanceHandler{},
		WalletBalanceHandler: rpcservice.WalletBalanceHandler{},
	}
}

func main() {

	mr := jsonrpc.NewMethodRepository()

	for _, s := range []Servicer{NewUserService(), NewBalanceService()} {
		for _, h := range s.Handlers() {
			fmt.Printf("%s\n", s.MethodName(h))
			mr.RegisterMethod(s.MethodName(h), h, h.Params(), h.Result())
		}
	}

	http.Handle("/jrpc", mr)
	http.HandleFunc("/jrpc/debug", mr.ServeDebug)

	if err := http.ListenAndServe(":8080", http.DefaultServeMux); err != nil {
		log.Fatalln(err)
	}
}
