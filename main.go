package main

import (
	"context"
	"log"
	"marketbill-messaging-service/config"
	"marketbill-messaging-service/handlers"
	"marketbill-messaging-service/models"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handle(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	config.ReadConfig()

	eventType := request.QueryStringParameters["event"]
	profile := os.Getenv("PROFILE")
	log.Print("PROFILE : ", profile)
	log.Print("Event-type : ", eventType)

	r := models.LambdaResponse{}
	switch eventType {
	case "ping":
		return handlers.HealthCheck(request)
	case "sms":
		return handlers.HandleDefaultSMS(request)
	case "messaging":
		return r.Error(http.StatusInternalServerError, "no functions in this event...")
	default:
		return r.Error(http.StatusBadRequest, "Wrong event type")
	}
}

func main() {
	lambda.Start(handle)
}
