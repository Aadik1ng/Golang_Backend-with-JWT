package api

import (
	"daily-expenses/internal/user"

	"github.com/gorilla/mux"
)

func RegisterUserRoutes(r *mux.Router) {
	r.HandleFunc("/users", user.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", user.GetUser).Methods("GET")
}
