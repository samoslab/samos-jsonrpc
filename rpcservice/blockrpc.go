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
		Seq  string `json:"seq"`
		Hash string `json:"hash"`
	}
	BlockLastNRequest struct {
		Num string `json:"num"`
	}

	BlockRangeRequest struct {
		Start string `json:"start"`
		End   string `json:"end"`
	}

	BlockHandler struct {
		BackendServer string
	}
	BlockRangeHandler struct {
		BackendServer string
	}
	BlockLastNHandler struct {
		BackendServer string
	}
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

	if req.Hash == "" && req.Seq == "" {
		return nil, jsonrpc.ErrInvalidParams()
	}

	url := fmt.Sprintf("%s/block", h.BackendServer)
	if req.Hash != "" {
		url = fmt.Sprintf("%s?hash=%s", url, req.Hash)
	} else if req.Seq != "" {
		url = fmt.Sprintf("%s?seq=%s", url, req.Seq)
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

func (h BlockRangeHandler) Name() string {
	return "blockRange"
}

func (h BlockRangeHandler) Params() interface{} {
	return BlockRangeRequest{}
}

func (h BlockRangeHandler) Result() interface{} {
	return visor.ReadableBlocks{}
}

func (h BlockRangeHandler) ServeJSONRPC(c context.Context, params *fastjson.RawMessage) (interface{}, *jsonrpc.Error) {
	var req BlockRangeRequest
	if err := jsonrpc.Unmarshal(params, &req); err != nil {
		return nil, ErrCustomise(err)
	}

	if req.Start == "" || req.End == "" {
		return nil, jsonrpc.ErrInvalidParams()
	}

	url := fmt.Sprintf("%s/blocks?start=%s&end=%s", h.BackendServer, req.Start, req.End)
	fmt.Printf("url %s\n", url)
	byteBody, err := SendRequest("GET", url, nil)
	if err != nil {
		return nil, ErrCustomise(err)
	}
	output := visor.ReadableBlocks{}
	if err := json.Unmarshal(byteBody, &output); err != nil {
		return nil, ErrCustomise(err)
	}
	return output, nil
}

func (h BlockLastNHandler) Name() string {
	return "blockLastN"
}

func (h BlockLastNHandler) Params() interface{} {
	return BlockLastNRequest{}
}

func (h BlockLastNHandler) Result() interface{} {
	return visor.ReadableBlocks{}
}

func (h BlockLastNHandler) ServeJSONRPC(c context.Context, params *fastjson.RawMessage) (interface{}, *jsonrpc.Error) {
	var req BlockLastNRequest
	if err := jsonrpc.Unmarshal(params, &req); err != nil {
		return nil, ErrCustomise(err)
	}

	if req.Num == "" {
		return nil, jsonrpc.ErrInvalidParams()
	}

	url := fmt.Sprintf("%s/last_blocks?num=%s", h.BackendServer, req.Num)
	fmt.Printf("url %s\n", url)
	byteBody, err := SendRequest("GET", url, nil)
	if err != nil {
		return nil, ErrCustomise(err)
	}
	output := visor.ReadableBlocks{}
	if err := json.Unmarshal(byteBody, &output); err != nil {
		return nil, ErrCustomise(err)
	}
	return output, nil
}
