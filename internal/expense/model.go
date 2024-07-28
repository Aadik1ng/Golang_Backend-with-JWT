package expense

import (
	"time"

	"github.com/google/uuid"
)

type Participant struct {
	UserID     uuid.UUID `json:"userId"`
	Amount     float64   `json:"amount,omitempty"`     // For Exact split method
	Percentage float64   `json:"percentage,omitempty"` // For Percentage split method
}

type Expense struct {
	ID           uuid.UUID     `json:"id"`
	Description  string        `json:"description"`
	Amount       float64       `json:"amount"`
	SplitMethod  string        `json:"splitMethod"`
	Participants []Participant `json:"participants"`
	CreatedAt    time.Time     `json:"CreatedAt"`
}

type UserBalance struct {
	UserID  uuid.UUID `json:"userId"`
	Balance float64   `json:"balance"`
}

type BalanceSheet struct {
	Balances []UserBalance `json:"balances"`
}
