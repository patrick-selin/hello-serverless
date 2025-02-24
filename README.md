# Serverless TODO App with Go, AWS Lambda, and CDK (TypeScript)

This is a world simplest and unsafest serverless TODO application built with **Go** for the backend logic, deployed on **AWS Lambda**, and exposed through an **API Gateway (REST API)**. The infrastructure is provisioned using **AWS CDK (TypeScript)**.

## Overview

This project demonstrates how to:
- Build a serverless REST API with **GoLang**.
- Deploy a **Go Lambda function** on AWS.
- Expose the Lambda through **API Gateway** endpoints:
  - `GET /todos`: Retrieve a list of todos.
  - `POST /todos`: Create a new todo with a randomly generated number.
- Provision and manage all infrastructure using **AWS CDK (TypeScript)**.

## How It Works

- `GET /todos`: Returns a hardcoded list of todos.
- `POST /todos`: Accepts a JSON body to create a new todo with a randomly generated number.


## NOTES:

- Create Go binary: 
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bootstrap main.go   
- Set the Lambda Runtime settings -> handler: 
bootstrap


- Go jutut

A type assertion with an "ok pattern"
In Go, a type assertion allows you to extract the underlying value from an interface.

```
if v, ok := item["id"].(*types.AttributeValueMemberS); ok {
	todo.ID = v.Value
}
```

Group Variable Declaration Block

```
var (
	dynamoClient *dynamodb.Client
	tableName    = "TodosTable"
)

on sama kuin

var dynamoClient *dynamodb.Client
var tableName = "TodosTable"
```

Multiple Assignment in Go
```
variable1, variable2 := functionCall()
cfg, err := config.LoadDefaultConfig(context.TODO())

```

Error Handling Pattern:
```
if err != nil {
    // Handle the error (e.g., log it, return an error response, etc.)
    panic("unable to load SDK config, " + err.Error())
}
```