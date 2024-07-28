package expense

import (
	"errors"
)

func ValidateExpense(expense Expense) error {
	if expense.Description == "" {
		return errors.New("description is required")
	}

	if expense.Amount <= 0 {
		return errors.New("amount must be greater than 0")
	}

	if len(expense.Participants) == 0 {
		return errors.New("at least one participant is required")
	}

	switch expense.SplitMethod {
	case "equal":
		// no additional validation needed
	case "exact":
		total := 0.0
		for _, p := range expense.Participants {
			total += p.Amount
		}
		if total != expense.Amount {
			return errors.New("total exact amounts do not match the expense amount")
		}
	case "percentage":
		total := 0.0
		for _, p := range expense.Participants {
			total += p.Percentage
		}
		if total != 100 {
			return errors.New("total percentages must add up to 100")
		}
	default:
		return errors.New("invalid split method")
	}

	return nil
}
