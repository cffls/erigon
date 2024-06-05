package da

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gateway-fm/cdk-erigon-lib/common"
	"github.com/ledgerwatch/erigon/common/hexutil"
)

const maxAttempts = 10
const retryDelay = 500 * time.Millisecond

// ErrorObject is a jsonrpc error
type ErrorObject struct {
	Code    int     `json:"code"`
	Message string  `json:"message"`
	Data    *[]byte `json:"data,omitempty"`
}

// Request is a jsonrpc request
type Request struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      interface{}     `json:"id"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params,omitempty"`
}

// Response is a jsonrpc  success response
type Response struct {
	JSONRPC string       `json:"jsonrpc"`
	ID      interface{}  `json:"id"`
	Result  string       `json:"result"`
	Error   *ErrorObject `json:"error"`
}

func JSONRPCCallWithContext(ctx context.Context, url, method string, parameters ...interface{}) (Response, error) {
	attemp := 0

	for attemp < maxAttempts {
		httpReq, err := BuildJsonHTTPRequest(ctx, url, method, parameters...)
		if err != nil {
			return Response{}, err
		}

		httpRes, err := http.DefaultClient.Do(httpReq)
		if err != nil {
			return Response{}, err
		}

		if httpRes.Body != nil {
			defer httpRes.Body.Close()
		}

		if httpRes.StatusCode == http.StatusTooManyRequests {
			time.Sleep(retryDelay)
			attemp += 1
			continue
		}

		if httpRes.StatusCode != http.StatusOK {
			return Response{}, fmt.Errorf("invalid status code, expected: %v, found: %v", http.StatusOK, httpRes.StatusCode)
		}

		var res Response
		if err = json.NewDecoder(httpRes.Body).Decode(&res); err != nil {
			return Response{}, err
		}

		return res, nil
	}

	return Response{}, fmt.Errorf("max attempts of data fetching reached, attemps: %v, DA url: %s", maxAttempts, url)
}

// BuildJsonHTTPRequest creates JSON RPC http request using provided url, method and parameters
func BuildJsonHTTPRequest(ctx context.Context, url, method string, parameters ...interface{}) (*http.Request, error) {
	params, err := json.Marshal(parameters)
	if err != nil {
		return nil, err
	}

	req := Request{
		JSONRPC: "2.0",
		ID:      float64(1),
		Method:  method,
		Params:  params,
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	return BuildJsonHttpRequestWithBody(ctx, url, reqBody)
}

// BuildJsonHttpRequestWithBody creates JSON RPC http request using provided url and request body
func BuildJsonHttpRequestWithBody(ctx context.Context, url string, reqBody []byte) (*http.Request, error) {
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(reqBody))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Add("Content-type", "application/json")

	return httpReq, nil
}

func GetOffChainData(ctx context.Context, url string, hash common.Hash) ([]byte, error) {
	response, err := JSONRPCCallWithContext(ctx, url, "sync_getOffChainData", hash)
	if err != nil {
		return nil, err
	}

	if response.Error != nil {
		return nil, fmt.Errorf("%v %v", response.Error.Code, response.Error.Message)
	}

	return hexutil.Decode(response.Result)
}
