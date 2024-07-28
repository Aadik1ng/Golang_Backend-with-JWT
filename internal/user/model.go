package user

import "github.com/google/uuid"

type User struct {
	ID     uuid.UUID `json:"id"`
	Email  string    `json:"email"`
	Name   string    `json:"name"`
	Mobile string    `json:"mobile"`
}
