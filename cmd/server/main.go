package main

import (
	"daily-expenses/api"
	"daily-expenses/auth"
	"daily-expenses/internal/expense"
	"daily-expenses/internal/user"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Register user routes
	api.RegisterUserRoutes(r)

	// Initialize sample data
	if err := initializeSampleData(); err != nil {
		log.Fatalf("Error initializing sample data: %v", err)
	}

	// Apply JWT Middleware to routes that need authentication
	apiRoutes := r.PathPrefix("/api").Subrouter()
	apiRoutes.Use(auth.AuthMiddleware) // Ensure middleware is used here

	// Register expense routes
	api.RegisterExpenseRoutes(apiRoutes)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func initializeSampleData() error {
	// Predefined UUIDs for consistent testing
	user1ID, err := uuid.Parse("1bc9e841-79f8-4050-88ea-76bd799326ae")
	if err != nil {
		return err
	}
	user2ID, err := uuid.Parse("2bc9e841-79f8-4050-88ea-76bd799326ae")
	if err != nil {
		return err
	}

	// Create users with fixed UUIDs
	if _, err := user.CreateUserService("user1@example.com", "John Doe", "1234567890", user1ID); err != nil {
		return err
	}
	if _, err := user.CreateUserService("user2@example.com", "Jane Doe", "0987654321", user2ID); err != nil {
		return err
	}

	// Create expenses with fixed UUIDs
	if _, err := expense.CreateExpenseService("Dinner", 50.0, "equal", []expense.Participant{
		{UserID: user1ID, Amount: 25.0},
		{UserID: user2ID, Amount: 25.0},
	}); err != nil {
		return err
	}
	if _, err := expense.CreateExpenseService("Taxi", 20.0, "equal", []expense.Participant{
		{UserID: user1ID, Amount: 10.0},
		{UserID: user2ID, Amount: 10.0},
	}); err != nil {
		return err
	}

	return nil
}
