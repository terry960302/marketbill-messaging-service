package handlers

import (
	"encoding/json"
	"fmt"
	"marketbill-messaging-service/constants"
	"marketbill-messaging-service/datastore"
	"marketbill-messaging-service/models"
	"marketbill-messaging-service/services"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func HandleSMS(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	r := models.NewLambdaResponse()

	req := models.MessagingRequest{}
	db, err := datastore.NewPostgresql()
	if err != nil {
		return r.Error(http.StatusInternalServerError, err.Error())
	}
	smsService := services.NewSmsService(db)

	if err := json.Unmarshal([]byte(request.Body), &req); err != nil {
		e := fmt.Sprintf("[HandleSMS] Unmarshal > %s", err.Error())
		return r.Error(http.StatusInternalServerError, e)
	}

	switch req.MessageType {
	case constants.Default.String():
		msg := req.Args[0].(string)
		res, err := smsService.SendDefaultSMS(req.To, msg, constants.SMS)
		if err != nil {
			return r.Error(http.StatusInternalServerError, err.Error())
		}
		return r.Json(http.StatusOK, res)
	case constants.Verification.String():
		res, err := smsService.SendSmsUsingTemplate(req.To, constants.Verification.Template(), 1, req.Args...)
		if err != nil {
			return r.Error(http.StatusInternalServerError, err.Error())
		}
		return r.Json(http.StatusOK, res)
	case constants.ApplyBizConnection.String():
		res, err := smsService.SendSmsUsingTemplate(req.To, constants.ApplyBizConnection.Template(), 2, req.Args...)
		if err != nil {
			return r.Error(http.StatusInternalServerError, err.Error())
		}
		return r.Json(http.StatusOK, res)
	case constants.ConfirmBizConnection.String():
		res, err := smsService.SendSmsUsingTemplate(req.To, constants.ConfirmBizConnection.Template(), 3, req.Args...)
		if err != nil {
			return r.Error(http.StatusInternalServerError, err.Error())
		}
		return r.Json(http.StatusOK, res)
	case constants.RejectBizConnection.String():
		res, err := smsService.SendSmsUsingTemplate(req.To, constants.RejectBizConnection.Template(), 2, req.Args...)
		if err != nil {
			return r.Error(http.StatusInternalServerError, err.Error())
		}
		return r.Json(http.StatusOK, res)
	case constants.IssueOrderSheetReceipt.String():
		res, err := smsService.SendSmsUsingTemplate(req.To, constants.IssueOrderSheetReceipt.Template(), 3, req.Args...)
		if err != nil {
			return r.Error(http.StatusInternalServerError, err.Error())
		}
		return r.Json(http.StatusOK, res)
	default:
		return r.Error(http.StatusBadRequest, "Bad Request : Wrong message type")
	}
}
