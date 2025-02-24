package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type Todo struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Num  int    `json:"num"`
}

var (
	dynamoClient *dynamodb.Client
	tableName    = "TodosTable"
)


func init() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}
	dynamoClient = dynamodb.NewFromConfig(cfg)
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
	out, err := dynamoClient.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	})
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: err.Error()}, nil
	}

	var todos []Todo
	for _, item := range out.Items {
		var todo Todo
	
		if v, ok := item["id"].(*types.AttributeValueMemberS); ok {
			todo.ID = v.Value
		}
	
		if v, ok := item["text"].(*types.AttributeValueMemberS); ok {
			todo.Text = v.Value
		}
	
		if v, ok := item["num"].(*types.AttributeValueMemberN); ok {
			num, err := strconv.Atoi(v.Value)
			if err == nil {
				todo.Num = num
			}
		}
	
		todos = append(todos, todo)
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
	
	todo.ID = fmt.Sprintf("%d", rand.Intn(1000000))
	todo.Num = rand.Intn(100)

	_, err = dynamoClient.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item: map[string]types.AttributeValue{
			"id":   &types.AttributeValueMemberS{Value: todo.ID},
			"text": &types.AttributeValueMemberS{Value: todo.Text},
			"num":  &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", todo.Num)},
		},
	})
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
			Headers: map[string]string{
				"Content-Type":                 "application/json",
				"Access-Control-Allow-Origin":  "*",
				"Access-Control-Allow-Methods": "GET, POST",
			},
		}, nil
	}

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
