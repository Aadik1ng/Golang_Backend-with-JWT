package user

import (
	"fmt"

	"github.com/google/uuid"
)

func CreateUserService(email, name, mobile string, userID uuid.UUID) User {
	user := User{
		ID:     userID,
		Email:  email,
		Name:   name,
		Mobile: mobile,
	}
	SaveUser(user)
	fmt.Println("User created with UUID:", user.ID)
	return user
}

func GetUserService(userID uuid.UUID) *User {
	return FetchUserByID(userID)
}
