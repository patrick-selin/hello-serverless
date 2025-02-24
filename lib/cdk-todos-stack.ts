import * as cdk from "aws-cdk-lib";
import { Construct } from "constructs";
import * as lambda from "aws-cdk-lib/aws-lambda";
import * as apigateway from "aws-cdk-lib/aws-apigateway";
import * as dynamodb from "aws-cdk-lib/aws-dynamodb";
import * as iam from "aws-cdk-lib/aws-iam";

import * as path from "path";

export class CdkTodosStack extends cdk.Stack {
  constructor(scope: Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    const todosLambda = new lambda.Function(this, "TodosFunction", {
      runtime: lambda.Runtime.PROVIDED_AL2,
      handler: "bootstrap",
      code: lambda.Code.fromAsset(path.join(__dirname, "../lambda/todos"), {
        bundling: {
          image: lambda.Runtime.PROVIDED_AL2.bundlingImage,
          command: [
            "bash",
            "-c",
            "GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o /asset-output/bootstrap main.go",
          ],
          volumes: [
            {
              hostPath: "/tmp/cdk-go-build",
              containerPath: "/go/pkg/mod",
            },
          ],
          environment: {
            GOCACHE: "/go/pkg/mod",
          },
        },
      }),
    });

    const todosTable = new dynamodb.Table(this, "TodosTable", {
      partitionKey: { name: "id", type: dynamodb.AttributeType.STRING },
      removalPolicy: cdk.RemovalPolicy.DESTROY, // NOT for production
    });

    todosTable.grantReadWriteData(todosLambda);
    todosLambda.addEnvironment("TABLE_NAME", todosTable.tableName);

    const api = new apigateway.RestApi(this, "TodosApi", {
      restApiName: "Todos Service",
      description: "This service serves todos.",
      defaultCorsPreflightOptions: {
        allowOrigins: apigateway.Cors.ALL_ORIGINS,
        allowMethods: ["GET", "POST"],
      },
    });

    const todos = api.root.addResource("todos");

    todos.addMethod(
      "GET",
      new apigateway.LambdaIntegration(todosLambda, {
        proxy: true,
      }),
      {
        methodResponses: [
          {
            statusCode: "200",
            responseParameters: {
              "method.response.header.Access-Control-Allow-Origin": true,
              "method.response.header.Content-Type": true,
            },
          },
        ],
      }
    );

    todos.addMethod(
      "POST",
      new apigateway.LambdaIntegration(todosLambda, {
        proxy: true,
      }),
      {
        methodResponses: [
          {
            statusCode: "201",
            responseParameters: {
              "method.response.header.Access-Control-Allow-Origin": true,
              "method.response.header.Content-Type": true,
            },
          },
        ],
      }
    );
  }
}
