package expense

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func AddExpense(w http.ResponseWriter, r *http.Request) {
	var expense Expense
	err := json.NewDecoder(r.Body).Decode(&expense)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdExpense, err := CreateExpenseService(expense.Description, expense.Amount, expense.SplitMethod, expense.Participants)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdExpense)
}

func HandleFetchAllExpenses(w http.ResponseWriter, r *http.Request) {
	allExpenses := FetchAllExpensesService()
	json.NewEncoder(w).Encode(allExpenses)
}

func HandleDownloadBalanceSheet(w http.ResponseWriter, r *http.Request) {
	balanceSheet := DownloadBalanceSheetService()
	w.Header().Set("Content-Disposition", "attachment; filename=balance_sheet.json")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(balanceSheet)
}

func GetUserExpenses(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := uuid.Parse(params["userId"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	expenses := GetUserExpensesService(userID)
	if expenses == nil {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(expenses)
}
