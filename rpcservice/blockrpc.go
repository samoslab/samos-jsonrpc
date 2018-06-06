package rpcservice

import (
	"context"
	"fmt"

	"github.com/intel-go/fastjson"
	"github.com/osamingo/jsonrpc"
	"github.com/samoslab/samos/src/gui"
	"github.com/skycoin/skycoin/src/visor"
)

type (
	BlockRequest struct {
		Seq  uint64 `json:"seq"`
		Hash string `json:"hash"`
	}
	BlockLastNRequest struct {
		Num int `json:"num"`
	}

	BlockRangeRequest struct {
		Start int `json:"start"`
		End   int `json:"end"`
	}

	BlockHandler struct {
		Client *gui.Client
	}
	BlockRangeHandler struct {
		Client *gui.Client
	}
	BlockLastNHandler struct {
		Client *gui.Client
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

	if req.Hash == "" && req.Seq == 0 {
		return nil, jsonrpc.ErrInvalidParams()
	}

	if req.Hash != "" {
		output, err := h.Client.BlockByHash(req.Hash)
		if err != nil {
			return nil, ErrCustomise(err)
		}
		return output, nil
	}

	output, err := h.Client.BlockBySeq(req.Seq)
	if err != nil {
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
	fmt.Printf("req:%+v\n", req)

	if req.Start == 0 || req.End == 0 {
		return nil, jsonrpc.ErrInvalidParams()
	}

	output, err := h.Client.Blocks(req.Start, req.End)
	if err != nil {
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

	if req.Num == 0 {
		return nil, jsonrpc.ErrInvalidParams()
	}

	output, err := h.Client.LastBlocks(req.Num)
	if err != nil {
		return nil, ErrCustomise(err)
	}
	return output, nil
}
