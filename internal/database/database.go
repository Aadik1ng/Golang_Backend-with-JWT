package database

import (
	"daily-expenses/internal/expense"
	"daily-expenses/internal/user"
	"sync"

	"github.com/google/uuid"
)

var Users = make(map[uuid.UUID]user.User)
var Expenses = make(map[uuid.UUID]expense.Expense)
var mu sync.Mutex
