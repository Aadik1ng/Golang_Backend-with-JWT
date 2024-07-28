package expense

import (
	"time"

	"github.com/google/uuid"
)

type Expense struct {
	ID           uuid.UUID     `json:"id"`
	Description  string        `json:"description"`
	Amount       float64       `json:"amount"`
	SplitMethod  string        `json:"splitMethod"`
	CreatedAt    time.Time     `json:"createdAt"`
	Participants []Participant `json:"participants"`
}

type Participant struct {
	UserID uuid.UUID `json:"userId"`
	Amount float64   `json:"amount"`
}
