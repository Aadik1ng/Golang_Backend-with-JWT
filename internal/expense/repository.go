package expense

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
)

var expenses = make(map[uuid.UUID]Expense)
var mu sync.Mutex

func SaveExpense(expense Expense) {
	mu.Lock()
	defer mu.Unlock()
	expenses[expense.ID] = expense
	fmt.Println("Saved expense:", expense) // Debug print
}

func FetchExpensesByUserID(userID uuid.UUID) []Expense {
	mu.Lock()
	defer mu.Unlock()
	var userExpenses []Expense
	for _, expense := range expenses {
		for _, participant := range expense.Participants {
			if participant.UserID == userID {
				userExpenses = append(userExpenses, expense)
				break
			}
		}
	}
	fmt.Println("Fetched expenses for user:", userID, userExpenses) // Debug print
	return userExpenses
}

func FetchAllExpenses() []Expense {
	mu.Lock()
	defer mu.Unlock()
	var allExpenses []Expense
	for _, expense := range expenses {
		allExpenses = append(allExpenses, expense)
	}
	fmt.Println("Fetched all expenses:", allExpenses) // Debug print
	return allExpenses
}
func FetchUserExpensesByID(userID uuid.UUID) []Expense {
	mu.Lock()
	defer mu.Unlock()
	var userExpenses []Expense
	if expense, ok := expenses[userID]; ok {
		userExpenses = append(userExpenses, expense)
	}
	fmt.Println("Fetched user expenses for user:", userID, userExpenses) // Debug print
	return userExpenses
}
