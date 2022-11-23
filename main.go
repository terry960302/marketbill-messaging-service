package main

import (
	"context"
	"log"
	"marketbill-messaging-service/handlers"
	"marketbill-messaging-service/models"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func init() {
	profile := os.Getenv("PROFILE")
	log.Print("PROFILE : ", profile)
}

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	r := models.LambdaResponse{}
	switch request.HTTPMethod {
	case "GET":
		return handlers.HealthCheck(request)
	case "POST":
		return handlers.HandleDefaultSMS(request)
	default:
		return r.Error(http.StatusBadRequest, "Wrong http method")
	}
}

func main() {
	lambda.Start(HandleRequest)
}
