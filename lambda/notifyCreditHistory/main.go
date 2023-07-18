package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/shikachii/credit-history/domain/model"
	"github.com/shikachii/credit-history/lib"
)

func Handler(ctx context.Context) ([]model.CreditHistory, error) {
	// 前日と当日の0:00のCreditHistoryを取得する
	// chs, err := lib.ScanBetweenTimestamp("credit-history", "1689174000", "1689260400")
	chs, err := lib.ScanBetweenTimestamp("credit-history", lib.Yesterday(), lib.Today())

	// 3日前から当日の0:00のCreditHistoryを取得する
	// chs, err := lib.ScanBetweenTimestamp("credit-history", "1689087600", "1689260400")

	if err != nil {
		return nil, fmt.Errorf("dynamoDB error: %w", err)
	}

	var todayAmount int64 = 0
	for _, ch := range *chs {
		amount, err := strconv.ParseInt(ch.Amount, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("strconv error: %w", err)
		}
		todayAmount += amount
	}

	chs, err = lib.ScanBetweenTimestamp("credit-history", lib.Yesterday() - (24 * 60 * 60), lib.Yesterday())
	var yesterdayAmount int64 = 0
	for _, ch := range *chs {
		amount, err := strconv.ParseInt(ch.Amount, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("strconv error: %w", err)
		}
		yesterdayAmount += amount
	}

	text := fmt.Sprintf("前日の利用額: %d円 (前々日比 %+d円)", todayAmount, todayAmount - yesterdayAmount); 
	// \n前日の利用額: %s円(%s円)`, (*chs)[1].Amount)

	// slackに通知する
	sm := lib.SlackMessage{
		Channel:   "#credit-history",
		Username:  "通知くん",
		Text:      text,
		IconEmoji: ":green_heart:",
	}
	if err := lib.SendSlackMessage(sm); err != nil {
		return nil, fmt.Errorf("slack error: %w", err)
	}

	return *chs, nil
}

func main() {
	lambda.Start(Handler)
}