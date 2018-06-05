package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/osamingo/jsonrpc"
	"github.com/samoslab/samos/src/gui"
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

func NewHandlers(client *gui.Client) []HandleParamsResulter {
	return []HandleParamsResulter{
		rpcservice.BalanceHandler{Client: client},
		rpcservice.WalletBalanceHandler{Client: client},
		rpcservice.WalletSpentHandler{Client: client},
		rpcservice.WalletCreateHandler{Client: client},
		rpcservice.WalletHandler{Client: client},
		rpcservice.WalletEncryptHandler{Client: client},
		rpcservice.WalletDecryptHandler{Client: client},

		rpcservice.VersionHandler{Client: client},
		rpcservice.AddressNewHandler{Client: client},
		rpcservice.OutputsHandler{Client: client},
		rpcservice.BlockHandler{Client: client},
		rpcservice.BlockRangeHandler{Client: client},
		rpcservice.BlockLastNHandler{Client: client},
		rpcservice.TransactionHandler{Client: client},
		rpcservice.CreateTransactionHandler{Client: client},
		rpcservice.InjectTransactionHandler{Client: client},
	}
}

func main() {

	var nodeRpcAddress string
	var ListenAddr string
	flag.StringVar(&nodeRpcAddress, "backend", "http://127.0.0.1:8640", "backend server web interface addr")
	flag.StringVar(&ListenAddr, "port", "127.0.0.1:8081", "listen port")
	flag.Parse()
	fmt.Printf("backend addr %s\n", nodeRpcAddress)
	fmt.Printf("listen addr %s\n", ListenAddr)
	mr := jsonrpc.NewMethodRepository()

	client := gui.NewClient(nodeRpcAddress)
	if client == nil {
		fmt.Printf("connect to samos service failed")
		return
	}
	for _, h := range NewHandlers(client) {
		fmt.Printf("%s\n", h.Name())
		mr.RegisterMethod(h.Name(), h, h.Params(), h.Result())
	}

	http.Handle("/jrpc", mr)
	http.HandleFunc("/jrpc/debug", mr.ServeDebug)

	if err := http.ListenAndServe(ListenAddr, http.DefaultServeMux); err != nil {
		log.Fatalln(err)
	}
}
