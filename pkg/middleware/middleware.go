package middleware

import (
	"net/http"
)

// ExampleMiddleware is a simple middleware function
func ExampleMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do some work before the request
		// Example: logging, authentication, etc.
		// log.Println("Request received")

		// Pass the request to the next handler
		next.ServeHTTP(w, r)

		// Do some work after the request
		// Example: logging, metrics, etc.
		// log.Println("Request processed")
	})
}
