package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	// Create a gin router with default middleware
	r := gin.Default()

	// Define the GET /ping route
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	return r
}

func TestPingRoute(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Setup the router
	router := setupRouter()

	// Create a test HTTP recorder
	w := httptest.NewRecorder()

	// Create a test request to the /ping endpoint
	req, _ := http.NewRequest("GET", "/ping", nil)

	// Serve the request using the router
	router.ServeHTTP(w, req)

	// Assert that the response code is 200 OK
	assert.Equal(t, http.StatusOK, w.Code)

	// Assert that the response body is "pong"
	assert.Equal(t, "pong", w.Body.String())
}