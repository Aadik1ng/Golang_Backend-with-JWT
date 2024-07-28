package expense

import (
	"sync"

	"github.com/google/uuid"
)

var (
	expenses   = make(map[uuid.UUID]Expense)
	balances   = make(map[uuid.UUID]float64)
	expenseMux sync.Mutex
)

func SaveExpense(expense Expense) {
	expenseMux.Lock()
	defer expenseMux.Unlock()
	expenses[expense.ID] = expense
	for _, participant := range expense.Participants {
		balances[participant.UserID] -= participant.Amount
	}
}

func FetchAllExpenses() []Expense {
	expenseMux.Lock()
	defer expenseMux.Unlock()

	var allExpenses []Expense
	for _, exp := range expenses {
		allExpenses = append(allExpenses, exp)
	}
	return allExpenses
}

func FetchUserExpenses(userID uuid.UUID) []Expense {
	expenseMux.Lock()
	defer expenseMux.Unlock()

	var userExpenses []Expense
	for _, exp := range expenses {
		for _, participant := range exp.Participants {
			if participant.UserID == userID {
				userExpenses = append(userExpenses, exp)
				break
			}
		}
	}
	return userExpenses
}
