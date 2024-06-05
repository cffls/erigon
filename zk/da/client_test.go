package da

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gateway-fm/cdk-erigon-lib/common"
	"github.com/stretchr/testify/require"
)

func TestClient_GetOffChainData(t *testing.T) {
	tests := []struct {
		name       string
		hash       common.Hash
		result     string
		data       []byte
		statusCode int
		err        error
	}{
		{
			name:   "successfully got offhcain data",
			hash:   common.BytesToHash([]byte("hash")),
			result: fmt.Sprintf(`{"result":"0x%s"}`, hex.EncodeToString([]byte("offchaindata"))),
			data:   []byte("offchaindata"),
		},
		{
			name:   "error returned by server",
			hash:   common.BytesToHash([]byte("hash")),
			result: `{"error":{"code":123,"message":"test error"}}`,
			err:    errors.New("123 test error"),
		},
		{
			name:   "invalid offchain data returned by server",
			hash:   common.BytesToHash([]byte("hash")),
			result: `{"result":"invalid-signature"}`,
			err:    errors.New("hex string without 0x prefix"),
		},
		{
			name:       "unsuccessful status code returned by server",
			hash:       common.BytesToHash([]byte("hash")),
			statusCode: http.StatusUnauthorized,
			err:        errors.New("invalid status code, expected: 200, found: 401"),
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				var res Request
				require.NoError(t, json.NewDecoder(r.Body).Decode(&res))
				require.Equal(t, "sync_getOffChainData", res.Method)

				var params []common.Hash
				require.NoError(t, json.Unmarshal(res.Params, &params))
				require.Equal(t, tt.hash, params[0])

				if tt.statusCode > 0 {
					w.WriteHeader(tt.statusCode)
				}

				_, err := fmt.Fprint(w, tt.result)
				require.NoError(t, err)
			}))
			defer svr.Close()

			got, err := GetOffChainData(context.Background(), svr.URL, tt.hash)
			if tt.err != nil {
				require.Error(t, err)
				require.EqualError(t, tt.err, err.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.data, got)
			}
		})
	}
}

func TestJSONRPCCallWithContext(t *testing.T) {
	tests := []struct {
		name         string
		prepare      func(*httptest.Server)
		expectedResp Response
		expectedErr  string
		statusCode   int
	}{
		{
			name: "successful JSON-RPC call",
			prepare: func(svr *httptest.Server) {
				svr.Config.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					_, err := fmt.Fprint(w, `{"jsonrpc":"2.0","id":1,"result":"success"}`)
					require.NoError(t, err)
				})
			},
			expectedResp: Response{JSONRPC: "2.0", ID: float64(1), Result: "success"},
		},
		{
			name: "handle retry on 429",
			prepare: func(svr *httptest.Server) {
				callCount := 0
				svr.Config.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					callCount++
					if callCount <= 3 {
						w.WriteHeader(http.StatusTooManyRequests)
					} else {
						_, err := fmt.Fprint(w, `{"jsonrpc":"2.0","id":1,"result":"after retry"}`)
						require.NoError(t, err)
					}
				})
			},
			expectedResp: Response{JSONRPC: "2.0", ID: float64(1), Result: "after retry"},
		},
		{
			name: "handle too many retry on 429",
			prepare: func(svr *httptest.Server) {
				callCount := 0
				svr.Config.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					callCount++
					if callCount <= 100 {
						w.WriteHeader(http.StatusTooManyRequests)
					} else {
						_, err := fmt.Fprint(w, `{"jsonrpc":"2.0","id":1,"result":"after retry"}`)
						require.NoError(t, err)
					}
				})
			},
			expectedErr: "max attempts of data fetching reached",
		},
		{
			name:        "returns error on HTTP status not OK",
			statusCode:  http.StatusForbidden,
			expectedErr: "invalid status code, expected: 200, found: 403",
		},
		{
			name: "error decoding JSON response",
			prepare: func(svr *httptest.Server) {
				svr.Config.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					_, err := fmt.Fprint(w, `{"jsonrpc":"2.0","id":1,"result":`)
					require.NoError(t, err)
				})
			},
			expectedErr: "unexpected EOF",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svr := httptest.NewServer(nil)
			defer svr.Close()

			if tt.prepare != nil {
				tt.prepare(svr)
			}

			if tt.statusCode != 0 {
				svr.Config.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(tt.statusCode)
				})
			}

			resp, err := JSONRPCCallWithContext(context.Background(), svr.URL, "testMethod")
			if tt.expectedErr != "" {
				require.Error(t, err)
				require.Contains(t, err.Error(), tt.expectedErr)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expectedResp, resp)
			}
		})
	}
}
