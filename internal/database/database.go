package database

import (
	"daily-expenses/internal/expense"
	"daily-expenses/internal/user"
	"sync"

	"github.com/google/uuid"
)

var (
	Users    = make(map[uuid.UUID]user.User)
	Expenses = make(map[uuid.UUID]expense.Expense)
	mu       sync.Mutex
)

func GetUser(userID uuid.UUID) (user.User, bool) {
	mu.Lock()
	defer mu.Unlock()
	user, exists := Users[userID]
	return user, exists
}

func AddUser(user user.User) {
	mu.Lock()
	defer mu.Unlock()
	Users[user.ID] = user
}

func GetExpense(expenseID uuid.UUID) (expense.Expense, bool) {
	mu.Lock()
	defer mu.Unlock()
	exp, exists := Expenses[expenseID]
	return exp, exists
}

func AddExpense(exp expense.Expense) {
	mu.Lock()
	defer mu.Unlock()
	Expenses[exp.ID] = exp
}
