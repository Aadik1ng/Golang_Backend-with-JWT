package expense

import (
	"errors"
	"strconv"
	"time"

	"github.com/google/uuid"
)

func CreateExpenseService(description string, amount float64, splitMethod string, participants []Participant) (Expense, error) {
	expense := Expense{
		ID:           uuid.New(),
		Description:  description,
		Amount:       amount,
		SplitMethod:  splitMethod,
		Participants: participants,
		CreatedAt:    time.Now(),
	}

	if numericSplit, err := strconv.Atoi(splitMethod); err == nil {
		if numericSplit < 0 || numericSplit > 100 {
			return Expense{}, errors.New("invalid percentage value")
		}
		splitPercentageBasedOnValue(&expense, float64(numericSplit))
	} else {
		switch splitMethod {
		case "equal":
			splitEqual(&expense)
		case "exact":
			if err := splitExact(&expense); err != nil {
				return Expense{}, err
			}
		case "percentage":
			if err := splitPercentage(&expense); err != nil {
				return Expense{}, err
			}
		default:
			return Expense{}, errors.New("invalid split method")
		}
	}

	saveExpense(expense)
	return expense, nil
}

func splitPercentageBasedOnValue(expense *Expense, percentage float64) {
	if len(expense.Participants) != 1 {
		expense.Participants[0].Percentage = 100 - percentage
		expense.Participants = append(expense.Participants, Participant{
			UserID:     uuid.New(),
			Percentage: percentage,
		})
	} else {
		expense.Participants[0].Percentage = 100 - percentage
		expense.Participants = append(expense.Participants, Participant{
			UserID:     uuid.New(),
			Percentage: percentage,
		})
	}
	for i := range expense.Participants {
		expense.Participants[i].Amount = (expense.Amount * expense.Participants[i].Percentage) / 100
	}
}

func FetchAllExpensesService() []Expense {
	expenseMux.Lock()
	defer expenseMux.Unlock()

	var allExpenses []Expense
	for _, exp := range expenses {
		allExpenses = append(allExpenses, exp)
	}
	return allExpenses
}

func DownloadBalanceSheetService() BalanceSheet {
	expenseMux.Lock()
	defer expenseMux.Unlock()

	var balanceSheet BalanceSheet
	for userID, balance := range balances {
		balanceSheet.Balances = append(balanceSheet.Balances, UserBalance{
			UserID:  userID,
			Balance: balance,
		})
	}
	return balanceSheet
}

func GetUserExpensesService(userID uuid.UUID) []Expense {
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

func saveExpense(expense Expense) {
	expenseMux.Lock()
	defer expenseMux.Unlock()

	expenses[expense.ID] = expense
	for _, participant := range expense.Participants {
		balances[participant.UserID] -= participant.Amount
	}
}

func splitEqual(expense *Expense) {
	share := expense.Amount / float64(len(expense.Participants))
	for i := range expense.Participants {
		expense.Participants[i].Amount = share
	}
}

func splitExact(expense *Expense) error {
	total := 0.0
	for _, participant := range expense.Participants {
		total += participant.Amount
	}
	if total != expense.Amount {
		return errors.New("total amount does not match the sum of participants' amounts")
	}
	return nil
}

func splitPercentage(expense *Expense) error {
	total := 0.0
	for _, participant := range expense.Participants {
		total += participant.Percentage
	}
	if total != 100 {
		return errors.New("total percentage does not add up to 100")
	}

	for i := range expense.Participants {
		expense.Participants[i].Amount = (expense.Amount * expense.Participants[i].Percentage) / 100
	}
	return nil
}
