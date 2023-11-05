package login

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// LimitMiddleware is an HTTP middleware that rate limits requests.
func LimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		limiter := getUserLimiter(c.GetString("level"))
		// var limiter = rate.NewLimiter(1, 1)
		if !limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Too Many Requests"})
			return
		}
		// Process the next handler if the rate limit is not reached.
		c.Next()
	}
}

var (
	rateLimiters = make(map[string]*rate.Limiter)
	mu           sync.Mutex
)

func InitRateLimit() {
	// Assuming you have a list of user IDs
	mu.Lock()
	rateLimiters["1"] = rate.NewLimiter(1, 3)
	rateLimiters["2"] = rate.NewLimiter(1, 5)
	rateLimiters["3"] = rate.NewLimiter(1, 7)
	mu.Unlock()
}

// getUserLimiter gets or creates a rate limiter for the provided userID
func getUserLimiter(level string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	limiter, exists := rateLimiters[level]
	if !exists {
		// Create a new rate limiter and add it to the map
		limiter = rate.NewLimiter(1, 1) // 1 request/second with a burst of 1
		rateLimiters[level] = limiter
	}

	return limiter
}
