package recordCreditHistory

import (
	"context"
	"fmt"

	"github.com/shikachii/credit-history/domain/model"
	"github.com/shikachii/credit-history/lib"
)

type MyEvent struct {
	Email string `json:"email"`
}

// Handler is the entry point for the lambda function
func Handler(ctx context.Context, event MyEvent) ([]model.CreditHistory, error) {
	// ch, err := lib.Parse(event.Email)
	// if err != nil {
	// 	return *ch, err
	// }

	// output, err := lib.PutItem("credit-history", ch)
	chs, err := lib.ScanBetweenTimestamp("credit-history", "1689174000", "1689260400")
	if err != nil {
		return nil, fmt.Errorf("dynamoDB error: %w", err)
	}

	for _, ch := range *chs {
		fmt.Println(ch)
	}

	return *chs, nil
}