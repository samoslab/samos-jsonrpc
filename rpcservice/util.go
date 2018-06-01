package rpcservice

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/osamingo/jsonrpc"
)

func SendRequest(method, url string, reqBody []byte) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewReader(reqBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rsp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()
	byteBody, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	return byteBody, nil
}

// ErrCustomise returns invalid error by error.
func ErrCustomise(err error) *jsonrpc.Error {
	return &jsonrpc.Error{
		Code:    jsonrpc.ErrorCodeInternal,
		Message: err.Error(),
	}
}
