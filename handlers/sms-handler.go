package handlers

import (
	"encoding/json"
	"fmt"
	"marketbill-messaging-service/constants"
	"marketbill-messaging-service/models"
	"marketbill-messaging-service/services"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func HandleDefaultSMS(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	r := models.NewLambdaResponse()
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

func HandleSMS(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	r := models.NewLambdaResponse()

	req := models.MessagingRequest{}

	if err := json.Unmarshal([]byte(request.Body), &req); err != nil {
		e := fmt.Sprintf("[HandleSMS] Unmarshal > %s", err.Error())
		r.Error(http.StatusInternalServerError, e)
	}

	switch req.MessageType {
	case constants.Default.String():
		msg := req.Args[0].(string)
		res, err := services.SendDefaultSMS(req.To, msg)
		if err != nil {
			return r.Error(http.StatusInternalServerError, err.Error())

		}
		return r.Json(http.StatusOK, res)
	case constants.Verification.String():
		break
	case constants.ApplyBizConnection.String():
		break
	case constants.ConfirmBizConnection.String():
		break
	case constants.RejectBizConnection.String():
		break
	case constants.IssueOrderSheetReceipt.String():
		break

	}

	return r.Json(http.StatusOK, "")
}
