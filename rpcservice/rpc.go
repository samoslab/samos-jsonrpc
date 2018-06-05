package rpcservice

import (
	"context"

	"github.com/intel-go/fastjson"
	"github.com/osamingo/jsonrpc"
	"github.com/samoslab/samos/src/gui"
	"github.com/skycoin/skycoin/src/visor"
)

type (
	CoinServer struct {
		Server string
	}
	StatusHandler struct {
		BackendServer string
		Client        *gui.Client
	}
	VersionHandler struct {
		BackendServer string
		Client        *gui.Client
	}
	OutputsHandler struct {
		BackendServer string
		Client        *gui.Client
	}

	OutputsRequest struct {
		Addrs  []string `json:"addrs"`
		Hashes []string `json:"hashes"`
	}

	AddressNewRequest struct {
		ID       string `json:"id"`
		Num      int    `json:"num"`
		Password string `json:"password"`
	}

	AddressNewHandler struct {
		BackendServer string
		Client        *gui.Client
	}
)

func (h AddressNewHandler) Name() string {
	return "addressNew"
}

func (h AddressNewHandler) Params() interface{} {
	return AddressNewRequest{}
}

func (h AddressNewHandler) Result() interface{} {
	return AddrNewResult{}
}

type AddrNewResult struct {
	Addresses []string `json:"addresses"`
}

func (h AddressNewHandler) ServeJSONRPC(c context.Context, params *fastjson.RawMessage) (interface{}, *jsonrpc.Error) {
	var req AddressNewRequest
	if err := jsonrpc.Unmarshal(params, &req); err != nil {
		return nil, ErrCustomise(err)
	}

	if req.ID == "" {
		return nil, jsonrpc.ErrInvalidParams()
	}

	output, err := h.Client.NewWalletAddress(req.ID, req.Num, req.Password)
	if err != nil {
		return nil, ErrCustomise(err)
	}
	return output, nil
}

func (h OutputsHandler) Name() string {
	return "outputs"
}

func (h OutputsHandler) Params() interface{} {
	return OutputsRequest{}
}

func (h OutputsHandler) Result() interface{} {
	return visor.ReadableOutputSet{}
}
func (h OutputsHandler) ServeJSONRPC(c context.Context, params *fastjson.RawMessage) (interface{}, *jsonrpc.Error) {
	req := &OutputsRequest{}
	if err := jsonrpc.Unmarshal(params, req); err != nil {
		return nil, ErrCustomise(err)
	}

	if len(req.Addrs) == 0 && len(req.Hashes) == 0 {
		return nil, jsonrpc.ErrInvalidParams()
	}

	if len(req.Addrs) != 0 {
		output, err := h.Client.OutputsForAddresses(req.Addrs)
		if err != nil {
			return nil, ErrCustomise(err)
		}
		return output, nil
	}

	output, err := h.Client.OutputsForHashes(req.Hashes)
	if err != nil {
		return nil, ErrCustomise(err)
	}
	return output, nil
}

func (h VersionHandler) Name() string {
	return "version"
}

func (h VersionHandler) Params() interface{} {
	return struct{}{}
}

func (h VersionHandler) Result() interface{} {
	return visor.BuildInfo{}
}

func (h VersionHandler) ServeJSONRPC(c context.Context, params *fastjson.RawMessage) (interface{}, *jsonrpc.Error) {
	output, err := h.Client.Version()
	if err != nil {
		return nil, ErrCustomise(err)
	}

	return output, nil
}
