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
	CoinServer struct {
		Server string
	}
	StatusHandler struct {
		BackendServer string
	}
	VersionHandler struct {
		BackendServer string
	}
	OutputsHandler struct {
		BackendServer string
	}

	OutputsRequest struct {
		Addrs  string `json:"addrs"`
		Hashes string `json:"hashes"`
	}

	AddressNewRequest struct {
		ID       string `json:"id"`
		Num      string `json:"num"`
		Password string `json:"password"`
	}

	AddressNewHandler struct {
		BackendServer string
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

	url := fmt.Sprintf("%s/wallet/newAddress", h.BackendServer)
	fmt.Printf("url %s\n", url)
	reqBody := fmt.Sprintf("id=%s&num=%s&password=%s", req.ID, req.Num, req.Password)
	byteBody, err := SendRequest("POST", url, []byte(reqBody))
	if err != nil {
		return nil, ErrCustomise(err)
	}
	output := AddrNewResult{}
	if err := json.Unmarshal(byteBody, &output); err != nil {
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

	if req.Addrs != "" && req.Hashes != "" {
		return nil, jsonrpc.ErrInvalidParams()
	}

	url := fmt.Sprintf("%s/outputs", h.BackendServer)
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
	url := fmt.Sprintf("%s/version", h.BackendServer)
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
