package main

import (
	"fmt"
	"sync"
	"time"
)

// This Rate limiter is implemented using "Token Bucket".
/*
There are other algorithms are available to do the task :
- Fixed Window
- Sliding Window
- Token Bucket (Implementation below)
- Leaky Bucket

*/

type RateLimiter struct {
	capacity       int           // Maximum number of tokens
	refillRate     int           // Tokens added per interval
	tokens         int           // Current available tokens
	lastRefill     time.Time     // Last refill timestamp
	refillInterval time.Duration // Refill interval (e.g., 1 second)
	mutex          sync.Mutex    // To protect concurrent access
}

func NewRateLimiter(capacity, refillRate int, refillInterval time.Duration) *RateLimiter {
	return &RateLimiter{
		capacity:       capacity,
		refillRate:     refillRate,
		tokens:         capacity, // intially it's capacity
		lastRefill:     time.Now(),
		refillInterval: refillInterval,
	}
}

func (rl *RateLimiter) refill() {
	now := time.Now()
	elapsed := now.Sub(rl.lastRefill)
	fmt.Println("Elapsed time:", elapsed)

	if elapsed >= rl.refillInterval {
		println("refilling the token")
		refillTokens := int(elapsed/rl.refillInterval) * rl.refillRate
		rl.tokens = min(rl.capacity, rl.tokens+refillTokens)
		rl.lastRefill = now
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (rl *RateLimiter) AllowRequest() bool {

	rl.mutex.Lock()
	defer rl.mutex.Unlock()

	rl.refill()

	if rl.tokens > 0 {
		rl.tokens--
		return true
	}
	return false
}
