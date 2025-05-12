package handlers

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/asutosh2203/url-shortener-go/storage"
	"github.com/asutosh2203/url-shortener-go/utils"
	"github.com/gin-gonic/gin"

	"crypto/sha256"
	"encoding/binary"
	"time"
)

// expected request body struct
type RequestBody struct {
	URL string `json:"url" binding:"required"`
	TTL int    `json:"ttl"` // in hours
}

func ShortenURL(ctx *gin.Context) {

	var req RequestBody

	// Bind the request body  with req JSON
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Validate the input URL
	if !utils.IsValidURL(req.URL) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input URL"})
		return
	}

	// Generate the short code for the URL
	shortUrlCode := generateRandomString(5, req.URL)

	expr := time.Duration(0)

	if req.TTL > 0 {
		expr = time.Duration(req.TTL) * time.Hour
	}

	// Set the key value pair in Redis
	if err := storage.Set(shortUrlCode, req.URL, expr); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store URL"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":  "URL shortened successfully",
		"shortUrl": "http://localhost:8080/" + shortUrlCode,
	})

}

func generateRandomString(length int, seedStr string) string {
	const charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// Mix time and string seed into a single int64 seed using SHA-256
	hash := sha256.Sum256(fmt.Appendf(nil, "%s%d", seedStr, time.Now().UnixNano()))
	seed := int64(binary.BigEndian.Uint64(hash[:8]))

	// Create a new Rand instance with the combined seed
	seededRand := rand.New(rand.NewSource(seed))

	code := make([]byte, length)
	for i := range code {
		code[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(code)
}
