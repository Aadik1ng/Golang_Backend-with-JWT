package user

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

func CreateUserService(email, name, mobile string, userID uuid.UUID) (User, error) {
	if email == "" || name == "" || mobile == "" {
		return User{}, errors.New("all fields are required")
	}

	user := User{
		ID:     userID,
		Email:  email,
		Name:   name,
		Mobile: mobile,
	}
	SaveUser(user)
	fmt.Println("User created with UUID:", user.ID)
	return user, nil
}

func GetUserService(userID uuid.UUID) *User {
	return FetchUserByID(userID)
}
