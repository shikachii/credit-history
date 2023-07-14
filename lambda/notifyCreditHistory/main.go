package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/shikachii/credit-history/domain/model"
	"github.com/shikachii/credit-history/lib"
)

func Handler(ctx context.Context) ([]model.CreditHistory, error) {
	
	chs, err := lib.ScanBetweenTimestamp("credit-history", "1689174000", "1689260400")
	if err != nil {
		return nil, fmt.Errorf("dynamoDB error: %w", err)
	}

	for _, ch := range *chs {
		fmt.Println(ch)
	}

	return *chs, nil
}

func main() {
	lambda.Start(Handler)
}