package api

import (
	"daily-expenses/internal/expense"

	"github.com/gorilla/mux"
)

func RegisterExpenseRoutes(r *mux.Router) {
	r.HandleFunc("/expenses", expense.AddExpense).Methods("POST")
	r.HandleFunc("/expenses/{userId}", expense.GetUserExpenses).Methods("GET")
	r.HandleFunc("/all-expenses", expense.HandleFetchAllExpenses).Methods("GET")
	r.HandleFunc("/balance-sheet", expense.HandleDownloadBalanceSheet).Methods("GET")
}
