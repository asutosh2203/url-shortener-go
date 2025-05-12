package middleware

import (
	"net/http"
	"strconv"
	"time"

	"github.com/asutosh2203/url-shortener-go/storage"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func RateLimiter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		rateVal, err := storage.Get(ctx.ClientIP())

		if err == redis.Nil {
			storage.Set(ctx.ClientIP(), "1", time.Duration(60)*time.Second)
		} else if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot get rate limit"})
			return
		} else {
			rateValue, err := strconv.Atoi(rateVal)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
				return
			}

			rateValue += 1

			if rateValue > 5 {
				ctx.JSON(http.StatusTooManyRequests, gin.H{"error": "You have exceed the number of requests"})
				return
			}

			ttlRemaining, err := storage.GetTTL(ctx.ClientIP())

			if err == nil {
				storage.Set(ctx.ClientIP(), strconv.Itoa(rateValue), ttlRemaining)
			}
		}

		ctx.Next()
	}
}
