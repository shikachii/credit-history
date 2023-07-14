package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/shikachii/credit-history/domain/model"
	"github.com/shikachii/credit-history/lib"
)

func Handler(ctx context.Context) ([]model.CreditHistory, error) {
	// 前日と当日の0:00のCreditHistoryを取得する
	chs, err := lib.ScanBetweenTimestamp("credit-history", "1689174000", "1689260400")

	// 3日前から当日の0:00のCreditHistoryを取得する
	// chs, err := lib.ScanBetweenTimestamp("credit-history", "1689087600", "1689260400")

	if err != nil {
		return nil, fmt.Errorf("dynamoDB error: %w", err)
	}

	for _, ch := range *chs {
		fmt.Println(ch)
	}

	// slackに通知する
	sm := lib.SlackMessage{
		Channel:   "#credit-history",
		Username:  "notifyCreditHistory",
		Text:      "test",
		IconEmoji: ":ghost:",
	}
	if err := lib.SendSlackMessage(sm); err != nil {
		return nil, fmt.Errorf("slack error: %w", err)
	}

	return *chs, nil
}

func main() {
	lambda.Start(Handler)
}