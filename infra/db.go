package infra

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type DynamoDB struct {
	*dynamodb.DynamoDB
}

func NewDynamoDB() *DynamoDB {
	mySession := session.Must(session.NewSession())
	var dynamoDB = dynamodb.New(mySession, aws.NewConfig().WithRegion("ap-northeast-1"))

	db := DynamoDB{dynamoDB}
	return &db
}

func (db *DynamoDB) PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	return db.DynamoDB.PutItem(input)
}
