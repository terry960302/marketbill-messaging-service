package handlers

import (
	"fmt"
	"marketbill-messaging-service/models"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
)

// func CheckServer(c echo.Context) error {
// 	profile := os.Getenv("PROFILE")
// 	msg := fmt.Sprintf("[%s]", profile) + "Marketbill Messaging Service is running..."
// 	return c.String(http.StatusOK, msg)
// }

func HealthCheck(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	r := models.LambdaResponse{}
	profile := os.Getenv("PROFILE")
	return r.Json(http.StatusOK, fmt.Sprintf("[%s] Marketbill messaging service is running...", profile))
}
