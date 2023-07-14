package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/shikachii/credit-history/domain/model"
	"github.com/shikachii/credit-history/lib"
)

type MyEvent struct {
	Email string `json:"email"`
}

// Handler is the entry point for the lambda function
func Handler(ctx context.Context, event MyEvent) (*model.CreditHistory, error) {
	ch, err := lib.Parse(event.Email)
	if err != nil {
		return ch, err
	}

	_, err = lib.PutItem("credit-history", ch)
	// chs, err := lib.ScanBetweenTimestamp("credit-history", "1689174000", "1689260400")
	if err != nil {
		return nil, fmt.Errorf("dynamoDB error: %w", err)
	}

	return ch, nil
}

func main() {
	lambda.Start(Handler)
}