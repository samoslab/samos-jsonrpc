package rpcservice

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/intel-go/fastjson"
	"github.com/osamingo/jsonrpc"
	"github.com/skycoin/skycoin/src/visor"
)

type (
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

	BlockHandler      struct{}
	BlockRangeHandler struct{}
	BlockLastNHandler struct{}

	TransactionHandler struct{}
)

func (h BlockHandler) Name() string {
	return "block"
}

func (h BlockHandler) Params() interface{} {
	return BlockRequest{}
}

func (h BlockHandler) Result() interface{} {
	return visor.ReadableBlock{}
}

func (h BlockHandler) ServeJSONRPC(c context.Context, params *fastjson.RawMessage) (interface{}, *jsonrpc.Error) {
	var req BlockRequest
	if err := jsonrpc.Unmarshal(params, &req); err != nil {
		return nil, ErrCustomise(err)
	}

	if req.Hash == "" {
		return nil, jsonrpc.ErrInvalidParams()
	}

	url := "http://127.0.0.1:6420/block"
	if req.Hash != "" {
		url = fmt.Sprintf("%s?hash=%s", url, req.Hash)
	} else if req.Seq != 0 {
		url = fmt.Sprintf("%s?seq=%d", url, req.Seq)
	}
	fmt.Printf("url %s\n", url)
	byteBody, err := SendRequest("GET", url, nil)
	if err != nil {
		return nil, ErrCustomise(err)
	}
	output := visor.ReadableBlock{}
	if err := json.Unmarshal(byteBody, &output); err != nil {
		return nil, ErrCustomise(err)
	}
	return output, nil
}
