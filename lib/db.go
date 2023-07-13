package lib

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
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
) (*dynamodb.ScanOutput, error) {
	// 7/13: 1689174000, 7/14: 1689260400
	// timestampの範囲を指定する
	filter := expression.Name("timestamp").Between(expression.Value(beforeTimestamp), expression.Value(afterTimestamp))
	expr, err := expression.NewBuilder().WithFilter(filter).Build()
	if err != nil {
		return nil, err
	}

	input := &dynamodb.ScanInput{
		TableName: &tableName,
		ExpressionAttributeNames: expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression: expr.Filter(),
		ProjectionExpression: expr.Projection(),
	}

	return dynamoDB.Scan(input)
}

// dynamoDBにデータを保存する
func PutItem(
	tableName string,
	item map[string]*dynamodb.AttributeValue,
) (*dynamodb.PutItemOutput, error) {
	input := &dynamodb.PutItemInput{
		TableName: &tableName,
		Item: item,
	}

	return dynamoDB.PutItem(input)
}