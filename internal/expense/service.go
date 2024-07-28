package expense

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func CreateExpenseService(description string, amount float64, splitMethod string, participants []Participant) Expense {
	expense := Expense{
		ID:           uuid.New(),
		Description:  description,
		Amount:       amount,
		SplitMethod:  splitMethod,
		CreatedAt:    time.Now(),
		Participants: participants,
	}
	SaveExpense(expense)

	// Print the UUID of the created expense
	fmt.Println("Expense created with UUID:", expense.ID)

	return expense
}

func GetExpensesByUserService(userID uuid.UUID) []Expense {
	return FetchExpensesByUserID(userID)
}

func FetchAllExpensesService() []Expense {
	return FetchAllExpenses()
}

func DownloadBalanceSheetService() []Expense {
	return FetchAllExpenses()
}

func GetUserExpensesService(userID uuid.UUID) []Expense {
	return FetchUserExpensesByID(userID)
}
