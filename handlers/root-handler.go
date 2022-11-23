package handlers

import (
	"fmt"
	"marketbill-messaging-service/models"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
)

func HealthCheck(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	r := models.LambdaResponse{}
	profile := os.Getenv("PROFILE")
	return r.Json(http.StatusOK, fmt.Sprintf("[%s] Marketbill messaging service is running....", profile))
}
