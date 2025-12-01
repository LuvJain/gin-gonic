package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a default gin router with Logger and Recovery middleware already attached
	router := gin.Default()

	// Define a GET route for /ping that returns "pong"
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Start the server on port 8080
	router.Run(":8080")
}