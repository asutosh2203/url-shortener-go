package handlers

import (
	"net/http"

	"github.com/asutosh2203/url-shortener-go/storage"
	"github.com/gin-gonic/gin"
)

func HandleRedirect(ctx *gin.Context) {
	code := ctx.Param("code")

	longUrl, err := storage.Get(code)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Failed to fetch URL"})
		return
	}

	ctx.Redirect(http.StatusFound, longUrl)
}
