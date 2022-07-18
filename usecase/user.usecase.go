package usecase

import (
	"encoding/json"
	"lambda_function_sample/domain/dynamodbmodel"
	"lambda_function_sample/interface/dynamodbrepository"

	"github.com/aws/aws-lambda-go/events"
)

type userUsecase struct {
	userDynamodbRepository dynamodbrepository.UserDynamoDBRepository
}

type UserUsecase interface {
	Create(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
}

func NewUserUsecase(userDynamodbRepository dynamodbrepository.UserDynamoDBRepository) UserUsecase {
	return &userUsecase{
		userDynamodbRepository: userDynamodbRepository,
	}
}

func (u *userUsecase) Create(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var item dynamodbmodel.UserItem
	if err := json.Unmarshal([]byte(request.Body), &item); err != nil {
		return events.APIGatewayProxyResponse{
			Body: err.Error(),
		}, err
	}
	input, err := u.userDynamodbRepository.ConvertPutItem(item)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body: err.Error(),
		}, nil
	}
	if _, err := u.userDynamodbRepository.PutItem(input); err != nil {
		return events.APIGatewayProxyResponse{
			Body: err.Error(),
		}, nil
	}
	b, err := json.Marshal(item)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
		}, err
	}
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(b),
	}, nil
}
