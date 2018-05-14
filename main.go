package main

import (
	"log"
	"net/http"

	"github.com/samoslab/skyjsonrpc/rpcservice"
)

func main() {
	mr := rpcservice.RegisterMethod()
	http.Handle("/jrpc", mr)
	http.HandleFunc("/jrpc/debug", mr.ServeDebug)

	if err := http.ListenAndServe(":8080", http.DefaultServeMux); err != nil {
		log.Fatalln(err)
	}
}
