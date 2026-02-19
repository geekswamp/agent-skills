package testutil

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

//
// ================================
// HTTP Request Builder
// ================================
//

// NewJSONRequest creates an HTTP request with JSON body.
func NewJSONRequest(t *testing.T, method, url string, body any) *http.Request {
	t.Helper()

	var buf bytes.Buffer

	if body != nil {
		RequireNoError(t, json.NewEncoder(&buf).Encode(body))
	}

	req := httptest.NewRequest(method, url, &buf)
	req.Header.Set("Content-Type", "application/json")

	return req
}

//
// ================================
// HTTP Recorder
// ================================
//

// NewRecorder returns a ResponseRecorder.
func NewRecorder() *httptest.ResponseRecorder {
	return httptest.NewRecorder()
}

//
// ================================
// Execute Handler
// ================================
//

// PerformRequest executes an HTTP handler.
func PerformRequest(handler http.Handler, req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	return rr
}

//
// ================================
// JSON Response Decoder
// ================================
//

// DecodeJSON decodes JSON response body.
func DecodeJSON[T any](t *testing.T, rr *httptest.ResponseRecorder) T {
	t.Helper()

	var result T
	RequireNoError(t, json.Unmarshal(rr.Body.Bytes(), &result))
	return result
}

//
// ================================
// Context Injection (optional)
// ================================
//

// WithContext injects context into request.
func WithContext(req *http.Request, ctx context.Context) *http.Request {
	return req.WithContext(ctx)
}
