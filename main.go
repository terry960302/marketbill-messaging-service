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

func init() {
	config.ReadConfig()
	profile := os.Getenv("PROFILE")
	log.Print("PROFILE : ", profile)
}

func LambdaHandler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	log.Print("[LambdaHandler] Hello world!")
	eventType := request.QueryStringParameters["event"]

	log.Print("Event-type : ", eventType)

	r := models.LambdaResponse{}

	switch eventType {
	case "ping":
		if request.HTTPMethod == "GET" {
			return handlers.HealthCheck(request)
		} else {
			return r.Error(http.StatusMethodNotAllowed, "Wrong http method")
		}
	case "sms":
		if request.HTTPMethod == "POST" {
			return handlers.HandleDefaultSMS(request)
		} else {
			return r.Error(http.StatusMethodNotAllowed, "Wrong http method")
		}

	case "messaging":
		return r.Error(http.StatusInternalServerError, "Still no functions in this event...")
	default:
		return r.Error(http.StatusBadRequest, "Wrong event type")
	}
}

func main() {
	lambda.Start(LambdaHandler)
}
