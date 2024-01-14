package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type MyEvent struct {
	UserID string `json:"user_id"`
}

type MyResponse struct {
	StatusCode int    `json:"status_code"`
	Body       string `json:"body"`
}

type Item struct {
	UserID    string `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type TableBasics struct {
	DynamoDbClient *dynamodb.Client
	TableName      string
}

func HandleRequest(ctx context.Context, event *MyEvent) (*MyResponse, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), func(o *config.LoadOptions) error {
		return nil
	})

	if err != nil {
		panic(err)
	}

	svc := dynamodb.NewFromConfig(cfg)

	out, err := svc.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName: aws.String("Users"),
	})

	if err != nil {
		panic(err)
	}

	for _, item := range out.Items {
		fmt.Println(item)
	}

	response := MyResponse{
		StatusCode: 200,
		Body:       "Hello, " + event.UserID,
	}

	return &response, nil
}

func main() {
	lambda.Start(HandleRequest)
}
