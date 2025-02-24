# Serverless TODO App with Go, AWS Lambda, and CDK (TypeScript)

This is a world simplest and unsafest serverless TODO application built with **Go** for the backend logic, deployed on **AWS Lambda**, and exposed through an **API Gateway (REST API)**. The infrastructure is provisioned using **AWS CDK (TypeScript)**.

![Image](https://github.com/user-attachments/assets/6e7e6757-ad4c-4a41-8600-fb9ddbe12a28)

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


## RANDOM NOTES:

- Create Go binary: 
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bootstrap main.go
- Zip for Lambda   
zip function.zip bootstrap
- Set the Lambda Runtime settings -> handler: 
bootstrap
- Uplod to AWS using CLI (kayta CDK jatakossa?)
aws lambda update-function-code \
  --function-name CdkTodosStack-TodosFunctionDE625C1E-2PjcztblBhYz \
  --zip-file fileb://function.zip
- Test the Lambda using CLI
aws lambda invoke \
  --function-name CdkTodosStack-TodosFunctionDE625C1E-2PjcztblBhYz \
  --payload '{}' \
  response.json


## GO LANGUAGE NOTES:
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


Iltaluettevaa:

AWS SDK GO v2
https://pkg.go.dev/github.com/aws/aws-sdk-go-v2

AWS CDK v2
https://docs.aws.amazon.com/cdk/api/v2/