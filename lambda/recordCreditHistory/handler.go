package recordCreditHistory

import (
	"context"

	"github.com/shikachii/credit-history/lib"
)

type MyEvent struct {
	Email string `json:"email"`
}

// Handler is the entry point for the lambda function
func Handler(ctx context.Context, event MyEvent) (lib.CreditHistory, error) {
	ch, err := lib.Parse(event.Email)
	if err != nil {
		return *ch, err
	}

	return *ch, nil
}