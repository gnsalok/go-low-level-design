package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

/*
**What:** add features without modifying core.
**Use for:** HTTP middleware: logging, auth, metrics.
---
* helloHandler is the core logic.
* withLogging is a decorator that wraps any http.Handler with extra behavior.
* You can easily chain more decorators like withAuth, withMetrics, etc.

*/

// --- Core handler (business logic) ---
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello!")
}

// --- Decorator: Logging middleware ---
func withLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Call the original handler
		next.ServeHTTP(w, r)

		// After handler completes, log info
		log.Printf("%s %s took %v", r.Method, r.URL.Path, time.Since(start))
	})
}

func main() {
	// Wrap helloHandler with logging decorator
	handler := withLogging(http.HandlerFunc(helloHandler))

	/*

	* same way you can write implementation for WithAPIKeyAuth, WithMetrics etc.

	 */

	http.Handle("/", handler)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
