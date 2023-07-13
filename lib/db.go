package lib

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"

	"github.com/google/uuid"
	"github.com/shikachii/credit-history/domain/model"
)

var dynamoDB = dynamodb.New(session.New(), aws.NewConfig().WithRegion("ap-northeast-1"))

func GetItem(
	tableName string,
	key map[string]*dynamodb.AttributeValue,
) (*dynamodb.GetItemOutput, error) {
	input := &dynamodb.GetItemInput{
		TableName: &tableName,
		Key: key,
	}

	return dynamoDB.GetItem(input)
}


// dynamoDBのtimestampフィールドから指定した期間のデータを取得する
func ScanBetweenTimestamp(
	tableName string,
	beforeTimestamp string,
	afterTimestamp string,
) (*[]model.CreditHistory, error) {
	// 7/13: 1689174000, 7/14: 1689260400
	// timestampの範囲を指定する
	filter := expression.Name("timestamp").Between(
		expression.Value((&dynamodb.AttributeValue{}).SetN(beforeTimestamp)),
		expression.Value((&dynamodb.AttributeValue{}).SetN(afterTimestamp)),
	)
	expr, err := expression.NewBuilder().WithFilter(filter).Build()
	if err != nil {
		return nil, err
	}

	input := &dynamodb.ScanInput{
		TableName: &tableName,
		ExpressionAttributeNames: expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression: expr.Filter(),
		// ProjectionExpression: expr.Projection(),
	}

	scan, err := dynamoDB.Scan(input)
	if err != nil {
		return nil, err
	}
	var ch []model.CreditHistory

	if err := dynamodbattribute.UnmarshalListOfMaps(scan.Items, &ch); err != nil {
		return nil, err
	}

	// ch.Date = *scan.Items[0]["date"].S
	// ch.Amount = *scan.Items[0]["amount"].N
	// ch.Card = *scan.Items[0]["card"].S
	// ch.Shop = *scan.Items[0]["shop"].S
	// ch.Timestamp = *scan.Items[0]["timestamp"].N
	// ch.Transaction = *scan.Items[0]["transaction"].S

	return &ch, err
}

// dynamoDBにデータを保存する
func PutItem(
	tableName string,
	ch *model.CreditHistory,
) (*dynamodb.PutItemOutput, error) {
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
	input := &dynamodb.PutItemInput{
		TableName: &tableName,
		Item: item,
	}

	return dynamoDB.PutItem(input)
}