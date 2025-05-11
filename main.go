package main

import (
	// "fmt"
	"net/http"

	"github.com/asutosh2203/url-shortener-go/handlers"
	"github.com/asutosh2203/url-shortener-go/storage"
	"github.com/gin-gonic/gin"
	// "time"
)

func main() {
	storage.InitRedis()
	r := gin.Default()

	// Home route
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Go URL shortener",
		})
	})

	// Endpoint to shorten URL
	r.POST("/shorten", handlers.ShortenURL)

	// Endpoint to redirect
	r.GET("/:code", handlers.HandleRedirect)

	// Start the server
	r.Run("0.0.0.0:8080")

}
