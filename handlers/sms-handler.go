package handlers

import (
	"encoding/json"
	"marketbill-messaging-service/models"
	"marketbill-messaging-service/services"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func HandleDefaultSMS(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
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
