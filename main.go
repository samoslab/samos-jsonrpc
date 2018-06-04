package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

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

func NewHandlers(NodeRpcAddress string) []HandleParamsResulter {
	return []HandleParamsResulter{
		rpcservice.BalanceHandler{NodeRpcAddress},
		rpcservice.WalletBalanceHandler{NodeRpcAddress},
		rpcservice.WalletSpentHandler{NodeRpcAddress},
		rpcservice.WalletCreateHandler{NodeRpcAddress},
		rpcservice.WalletHandler{NodeRpcAddress},
		rpcservice.WalletEncryptHandler{NodeRpcAddress},
		rpcservice.WalletDecryptHandler{NodeRpcAddress},

		rpcservice.VersionHandler{NodeRpcAddress},
		rpcservice.AddressNewHandler{NodeRpcAddress},
		rpcservice.OutputsHandler{NodeRpcAddress},
		rpcservice.BlockHandler{NodeRpcAddress},
		rpcservice.BlockRangeHandler{NodeRpcAddress},
		rpcservice.BlockLastNHandler{NodeRpcAddress},
		rpcservice.TransactionHandler{NodeRpcAddress},
		rpcservice.CreateTransactionHandler{NodeRpcAddress},
		rpcservice.InjectTransactionHandler{NodeRpcAddress},
	}
}

func main() {

	var NodeRpcAddress string
	var ListenAddr string
	flag.StringVar(&NodeRpcAddress, "backend", "http://127.0.0.1:8640", "backend server web interface addr")
	flag.StringVar(&ListenAddr, "port", "127.0.0.1:8081", "listen port")
	flag.Parse()
	fmt.Printf("backend addr %s\n", NodeRpcAddress)
	fmt.Printf("listen addr %s\n", ListenAddr)
	mr := jsonrpc.NewMethodRepository()

	for _, h := range NewHandlers(NodeRpcAddress) {
		fmt.Printf("%s\n", h.Name())
		mr.RegisterMethod(h.Name(), h, h.Params(), h.Result())
	}

	http.Handle("/jrpc", mr)
	http.HandleFunc("/jrpc/debug", mr.ServeDebug)

	if err := http.ListenAndServe(ListenAddr, http.DefaultServeMux); err != nil {
		log.Fatalln(err)
	}
}
