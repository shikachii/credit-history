package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/shikachii/credit-history/lambda/recordCreditHistory"
)

func main() {
	lambda.Start(recordCreditHistory.Handler)
}