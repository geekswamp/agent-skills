package testutil

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

// ================================
// Gin Setup
// ================================

// NewGinEngine creates a Gin engine in test mode.
func NewGinEngine() *gin.Engine {
	gin.SetMode(gin.TestMode)
	return gin.New()
}

// ================================
// Gin Perform Request
// ================================

// PerformGinRequest executes request against Gin engine.
func PerformGinRequest(
	t *testing.T,
	engine *gin.Engine,
	req *http.Request,
) *httptest.ResponseRecorder {
	t.Helper()

	recorder := httptest.NewRecorder()
	engine.ServeHTTP(recorder, req)
	return recorder
}

// ================================
// Gin Context
// ================================

// NewGinContext executes Gin with context.
func NewGinContext(t *testing.T) *gin.Context {
	t.Helper()
	gin.SetMode(gin.TestMode)

	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	RequireNotNil(t, ctx)

	return ctx
}

// NewGinContextWithRequest executes Gin with context and request.
func NewGinContextWithRequest(t *testing.T, req *http.Request) *gin.Context {
	t.Helper()
	gin.SetMode(gin.TestMode)

	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	
	RequireNotNil(t, ctx)
	RequireNotNil(t, req)

	ctx.Request = req
	return ctx
}
