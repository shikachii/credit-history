package recordCreditHistory

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/google/uuid"
	"github.com/shikachii/credit-history/domain/model"
	"github.com/shikachii/credit-history/lib"
)

type MyEvent struct {
	Email string `json:"email"`
}

// Handler is the entry point for the lambda function
func Handler(ctx context.Context, event MyEvent) (model.CreditHistory, error) {
	ch, err := lib.Parse(event.Email)
	if err != nil {
		return *ch, err
	}

	id := uuid.New().String()
	item := map[string]*dynamodb.AttributeValue{
		"id": {
			S: &id,
		},
		"date": {
			S: &ch.Date,
		},
		"shop": {
			S: &ch.Shop,
		},
		"amount": {
			N: &ch.Amount,
		},
		"transaction": {
			S: &ch.Transaction,
		},
		"card": {
			S: &ch.Card,
		},
		"timestamp": {
			N: &ch.Timestamp,
		},
	}
	output, err := lib.PutItem("credit-history", item)
	if err != nil {
		fmt.Println(output.String())
		return *ch, fmt.Errorf("dynamoDB error: %w", err)
		// return *ch, err
	}

	return *ch, nil
}