package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	php "github.com/deuill/go-php"
    "os"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
    engine, _ := php.New()
	
	context, _ := engine.NewContext()
    context.Output = os.Stdout
	
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       context.Exec("index.php"),
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}
