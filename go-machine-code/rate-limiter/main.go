package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	// Initialize the RateLimiter
	rl := NewRateLimiter(5, 2, 1*time.Second)

	// WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup
	wg.Add(10)

	// Start 10 concurrent goroutines
	for i := 1; i <= 10; i++ {
		go func(id int) {
			defer wg.Done()

			// Attempt to make 5 requests from each goroutine
			for j := 1; j <= 5; j++ {
				if rl.AllowRequest() {
					fmt.Printf("Goroutine %d: Request %d allowed.\n", id, j)
				} else {
					fmt.Printf("Goroutine %d: Request %d denied (rate limited).\n", id, j)
				}
				time.Sleep(400 * time.Millisecond) // Sleep for a short duration between requests
			}
		}(i)
	}

	// Wait for all goroutines to complete
	wg.Wait()

}
