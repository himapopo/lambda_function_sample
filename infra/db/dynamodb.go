package db

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func NewDynamoDB() (*dynamodb.DynamoDB, error) {
	// DB接続
	sess, err := session.NewSession()
	if err != nil {
		return nil, err
	}
	return dynamodb.New(sess), err
}
