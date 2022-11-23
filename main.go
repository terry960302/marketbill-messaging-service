package main

import (
	"context"
	"encoding/json"
	"log"
	"marketbill-messaging-service/models"
	"marketbill-messaging-service/services"
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
	req := models.DefaultSmsRequest{}

	if err := json.Unmarshal([]byte(request.Body), &req); err != nil {
		return r.Error(http.StatusBadRequest, err.Error())
	}

	res, err := services.SendDefaultSMS(req.To, req.Message)
	if err != nil {
		return r.Error(http.StatusInternalServerError, err.Error())

	}
	return r.Json(http.StatusOK, res)
}

func main() {
	lambda.Start(HandleRequest)
}
