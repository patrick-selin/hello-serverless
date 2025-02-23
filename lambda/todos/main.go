package main

import (
	"context"
	"encoding/json"
	"math/rand"

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
	case "POST":
		return createTodo(request)
	default:
		return events.APIGatewayProxyResponse{
			StatusCode: 405,
			Body:       `{"error": "Method Not Allowed"}`,
			Headers: map[string]string{
				"Content-Type":                 "application/json",
				"Access-Control-Allow-Origin":  "*",
				"Access-Control-Allow-Methods": "GET, POST",
			},
		}, nil
	}
}

func getTodos() (events.APIGatewayProxyResponse, error) {
	// Dummy data
	todos := []Todo{
		{ID: "1", Text: "Learn Go", Num: 42},
		{ID: "2", Text: "Build Serverless App", Num: 99},
	}
	body, _ := json.Marshal(todos)
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(body),
		Headers: map[string]string{
			"Content-Type":                 "application/json",
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Methods": "GET, POST",
		},
	}, nil
}

func createTodo(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var todo Todo
	err := json.Unmarshal([]byte(request.Body), &todo)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       `{"error": "Invalid request body"}`,
			Headers: map[string]string{
				"Content-Type":                 "application/json",
				"Access-Control-Allow-Origin":  "*",
				"Access-Control-Allow-Methods": "GET, POST",
			},
		}, nil
	}

	todo.ID = "random-id"
	todo.Num = rand.Intn(100)

	body, _ := json.Marshal(todo)
	return events.APIGatewayProxyResponse{
		StatusCode: 201,
		Body:       string(body),
		Headers: map[string]string{
			"Content-Type":                 "application/json",
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Methods": "GET, POST",
		},
	}, nil
}

func main() {
	lambda.Start(handleRequest)
}