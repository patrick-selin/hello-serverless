import * as cdk from "aws-cdk-lib";
import { Construct } from "constructs";
import * as lambda from "aws-cdk-lib/aws-lambda";
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
  }
}
