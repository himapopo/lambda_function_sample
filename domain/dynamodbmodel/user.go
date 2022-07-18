package dynamodbmodel

type UserItem struct {
	UserID    int    `json:"user_id" dynamodbav:"user_id"`
	LastName  string `json:"last_name" dynamodbav:"last_name"`
	FirstName string `json:"first_name" dynamodbav:"first_name"`
}
