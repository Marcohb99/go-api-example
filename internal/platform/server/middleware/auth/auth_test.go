package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMiddlewareFailureEmptyKey(t *testing.T) {
	// Setting up the Gin server
	err := os.Setenv("MHB_API_KEYS", "key1,key2,key3")
	gin.SetMode(gin.TestMode)
	engine := gin.New()
	engine.Use(Middleware())
	engine.GET(
		"/test-middleware",
		func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
		},
	)
	require.NoError(t, err)

	// Setting up the HTTP recorder and the request
	rec := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/test-middleware", nil)
	// Add custom headers to the request
	require.NoError(t, err)

	// WHEN: send request
	engine.ServeHTTP(rec, req)

	// THEN: get response and assert data
	res := rec.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusUnauthorized, res.StatusCode)
}

func TestMiddlewareFailureWrongKey(t *testing.T) {
	// Setting up the Gin server
	err := os.Setenv("MHB_API_KEYS", "key1,key2,key3")
	gin.SetMode(gin.TestMode)
	engine := gin.New()
	engine.Use(Middleware())
	engine.GET(
		"/test-middleware",
		func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
		},
	)
	require.NoError(t, err)

	// Setting up the HTTP recorder and the request
	rec := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/test-middleware", nil)
	// Add custom headers to the request
	req.Header.Set("X-API-KEY", "key89")

	require.NoError(t, err)

	// WHEN: send request
	engine.ServeHTTP(rec, req)

	// THEN: get response and assert data
	res := rec.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusUnauthorized, res.StatusCode)
}

func TestMiddlewareOk(t *testing.T) {
	// Setting up the Gin server
	err := os.Setenv("MHB_API_KEYS", "key1,key2,key3")
	gin.SetMode(gin.TestMode)
	engine := gin.New()
	engine.Use(Middleware())
	require.NoError(t, err)
	engine.GET(
		"/test-middleware",
		func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
		},
	)

	// Setting up the HTTP recorder and the request
	rec := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/test-middleware", nil)
	req.Header.Set("X-API-KEY", "key2")
	require.NoError(t, err)

	// WHEN: send request
	engine.ServeHTTP(rec, req)

	// THEN: get response and assert data
	res := rec.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusOK, res.StatusCode)
	data := rec.Body.Bytes()
	assert.Contains(t, string(data), "ok")
}
