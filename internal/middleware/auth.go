package middleware

import (
	"context"
	"net/http"
)

// Define a context key type to avoid context key collisions
type contextKey string

const ContextUserID contextKey = "userID"

// MiddlewareAuthenticate authenticates a user and adds their ID to the context
func MiddlewareAuthenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Example authentication logic
		userID := "1bc9e841-79f8-4050-88ea-76bd799326ae" // Replace with actual authentication logic
		ctx := context.WithValue(r.Context(), ContextUserID, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
