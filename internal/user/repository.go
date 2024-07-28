package user

import (
	"sync"

	"github.com/google/uuid"
)

var users = make(map[uuid.UUID]User)
var mu sync.Mutex

func SaveUser(user User) {
	mu.Lock()
	defer mu.Unlock()
	users[user.ID] = user
}

func FetchUserByID(userID uuid.UUID) *User {
	mu.Lock()
	defer mu.Unlock()
	if user, exists := users[userID]; exists {
		return &user
	}
	return nil
}
