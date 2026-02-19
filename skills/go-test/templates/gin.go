package testutil

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

//
// ================================
// Gin Setup
// ================================
//

// NewGinEngine creates a Gin engine in test mode.
func NewGinEngine() *gin.Engine {
	gin.SetMode(gin.TestMode)
	return gin.New()
}

//
// ================================
// Gin Perform Request
// ================================
//

// PerformGinRequest executes request against Gin engine.
func PerformGinRequest(
	t *testing.T,
	engine *gin.Engine,
	req *http.Request,
) *httptest.ResponseRecorder {
	t.Helper()

	rr := httptest.NewRecorder()
	engine.ServeHTTP(rr, req)
	return rr
}
