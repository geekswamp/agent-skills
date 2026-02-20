package testutil

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// ================================
// HTTP Request Builder
// ================================

// NewJSONRequest creates an HTTP request with JSON body.
func NewJSONRequest(t *testing.T, method, url, body string) *http.Request {
	t.Helper()

	req := httptest.NewRequest(method, url, strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	return req
}

// ================================
// HTTP Recorder
// ================================

// NewRecorder returns a ResponseRecorder.
func NewRecorder() *httptest.ResponseRecorder {
	return httptest.NewRecorder()
}

// ================================
// Execute Handler
// ================================

// PerformRequest executes an HTTP handler.
func PerformRequest(handler http.Handler, req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	return rr
}

// ================================
// JSON Response Decoder
// ================================

// DecodeJSON decodes JSON response body.
func DecodeJSON[T any](t *testing.T, rr *httptest.ResponseRecorder) T {
	t.Helper()

	var result T
	RequireNoError(t, json.Unmarshal(rr.Body.Bytes(), &result))
	return result
}

// ================================
// Context Injection (optional)
// ================================

// WithContext injects context into request.
func WithContext(req *http.Request, ctx context.Context) *http.Request {
	return req.WithContext(ctx)
}

// ================================
// HTTP Request
// ================================

// PostRequest creates an HTTP Post Request.
func PostRequest(target string, body io.Reader) *http.Request {
	return httptest.NewRequest(http.MethodPost, target, body)
}

// GetRequest creates an HTTP GET Request.
func GetRequest(target string, body io.Reader) *http.Request {
	return httptest.NewRequest(http.MethodGet, target, body)
}

// PatchRequest creates an HTTP PATCH Request.
func PatchRequest(target string, body io.Reader) *http.Request {
	return httptest.NewRequest(http.MethodPatch, target, body)
}

// DeleteRequest creates an HTTP DELETE Request.
func DeleteRequest(target string, body io.Reader) *http.Request {
	return httptest.NewRequest(http.MethodDelete, target, body)
}

// PutRequest creates an HTTP PUT Request.
func PutRequest(target string, body io.Reader) *http.Request {
	return httptest.NewRequest(http.MethodPut, target, body)
}