package recovery

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRecoveryMiddleware(t *testing.T)  {
	// Setup
	gin.SetMode(gin.TestMode)
	engine := gin.Default()
	engine.Use(Middleware())

	engine.GET("/test-middleware", func(ctx *gin.Context) {
		panic("test panic")
	})

	// Set up HTTP recorder and request
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/test-middleware", nil)
	require.NoError(t, err)

	// Assert request does not panic
	assert.NotPanics(t, func() {
		engine.ServeHTTP(w, req)
	})
}