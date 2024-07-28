package main

import (
	"daily-expenses/api"
	"daily-expenses/internal/expense"
	"daily-expenses/internal/user"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Register user and expense routes
	api.RegisterUserRoutes(r)
	api.RegisterExpenseRoutes(r)

	// Initialize sample data
	initializeSampleData()

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func initializeSampleData() {
	// Predefined UUIDs for consistent testing
	user1ID, _ := uuid.Parse("1bc9e841-79f8-4050-88ea-76bd799326ae")
	user2ID, _ := uuid.Parse("2bc9e841-79f8-4050-88ea-76bd799326ae")

	// Create users with fixed UUIDs
	user.CreateUserService("user1@example.com", "John Doe", "1234567890", user1ID)
	user.CreateUserService("user2@example.com", "Jane Doe", "0987654321", user2ID)

	// Create expenses with sample data
	expense.CreateExpenseService("Dinner", 50.0, "exact", []expense.Participant{
		{UserID: user1ID, Amount: 25.0},
		{UserID: user2ID, Amount: 25.0},
	})
	expense.CreateExpenseService("Taxi", 20.0, "exact", []expense.Participant{
		{UserID: user1ID, Amount: 10.0},
		{UserID: user2ID, Amount: 10.0},
	})
}
