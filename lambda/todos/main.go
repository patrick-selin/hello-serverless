package main

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Todo struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Num  int    `json:"num"`
}

func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch request.HTTPMethod {
	case "GET":
		return getTodos()
	default:
		return events.APIGatewayProxyResponse{StatusCode: 405}, nil
	}
}

func getTodos() (events.APIGatewayProxyResponse, error) {
	// Dummy data
	todos := []Todo{
		{ID: "1", Text: "Learn Go", Num: 42},
		{ID: "2", Text: "Build Serverless App", Num: 99},
	}
	body, _ := json.Marshal(todos)
	return events.APIGatewayProxyResponse{StatusCode: 200, Body: string(body)}, nil
}

func main() {
	lambda.Start(handleRequest)
}
