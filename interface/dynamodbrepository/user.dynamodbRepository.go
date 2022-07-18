package dynamodbrepository

import (
	"lambda_function_sample/domain/dynamodbmodel"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type userDynamoDBRepository struct {
	dynamoDB *dynamodb.DynamoDB
}

type UserDynamoDBRepository interface {
	PutItem(*dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error)
	ConvertPutItem(dynamodbmodel.UserItem) (*dynamodb.PutItemInput, error)
}

func NewUserDynamoDBRepository(db *dynamodb.DynamoDB) UserDynamoDBRepository {
	return &userDynamoDBRepository{
		dynamoDB: db,
	}
}

func (r *userDynamoDBRepository) PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	return r.dynamoDB.PutItem(input)
}

func (r *userDynamoDBRepository) ConvertPutItem(item dynamodbmodel.UserItem) (*dynamodb.PutItemInput, error) {
	inputAV, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		return nil, err
	}
	return &dynamodb.PutItemInput{
		TableName: aws.String("user"),
		Item:      inputAV,
	}, nil
}
