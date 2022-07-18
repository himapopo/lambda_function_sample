package main

import (
	"lambda_function_sample/infra/db"
	"lambda_function_sample/interface/dynamodbrepository"
	"lambda_function_sample/usecase"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Item struct {
	UserID    int    `json:"user_id" dynamodbav:"user_id"`
	LastName  string `json:"last_name" dynamodbav:"last_name"`
	FirstName string `json:"first_name" dynamodbav:"first_name"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	dynamoDB, err := db.NewDynamoDB()
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body: err.Error(),
		}, err
	}
	
	userRepository := dynamodbrepository.NewUserDynamoDBRepository(dynamoDB)
	userUsecase := usecase.NewUserUsecase(userRepository)
	return userUsecase.Create(request)
}

func main() {
	lambda.Start(handler)
}
