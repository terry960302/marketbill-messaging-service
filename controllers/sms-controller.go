package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"marketbill-messaging-service/constants"
	"marketbill-messaging-service/datastore"
	"marketbill-messaging-service/models"
	"marketbill-messaging-service/services"
	"net/http"
	"os"

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

	// dev 환경에서는 팀 내부에서만 SMS 보낼 수 있도록 제한
	if err := validatePhoneNoInDev(req.To); err != nil {
		msg := "Not supported phone number to send message(using SENS). Some phone numbers could restricted in local, dev environment. But not consider this issue as error."
		return r.Json(http.StatusAccepted, msg)
	}

	switch req.Template {
	case constants.Default.String():
		msg := req.Args[0].(string)
		res, err := smsService.SendDefaultSMS(req.To, msg, constants.SMS)
		if err != nil {
			return r.Error(http.StatusInternalServerError, err.Error())
		}
		return r.Json(http.StatusOK, res)
	case constants.Verification.String():
		res, err := smsService.SendSmsUsingTemplate(req.To, constants.Verification.Template(), constants.Verification.ArgsCount(), req.Args...)
		if err != nil {
			return r.Error(http.StatusInternalServerError, err.Error())
		}
		return r.Json(http.StatusOK, res)
	case constants.ApplyBizConnection.String():
		res, err := smsService.SendSmsUsingTemplate(req.To, constants.ApplyBizConnection.Template(), constants.ApplyBizConnection.ArgsCount(), req.Args...)
		if err != nil {
			return r.Error(http.StatusInternalServerError, err.Error())
		}
		return r.Json(http.StatusOK, res)
	case constants.ConfirmBizConnection.String():
		res, err := smsService.SendSmsUsingTemplate(req.To, constants.ConfirmBizConnection.Template(), constants.ConfirmBizConnection.ArgsCount(), req.Args...)
		if err != nil {
			return r.Error(http.StatusInternalServerError, err.Error())
		}
		return r.Json(http.StatusOK, res)
	case constants.RejectBizConnection.String():
		res, err := smsService.SendSmsUsingTemplate(req.To, constants.RejectBizConnection.Template(), constants.RejectBizConnection.ArgsCount(), req.Args...)
		if err != nil {
			return r.Error(http.StatusInternalServerError, err.Error())
		}
		return r.Json(http.StatusOK, res)
	case constants.IssueOrderSheetReceipt.String():
		res, err := smsService.SendSmsUsingTemplate(req.To, constants.IssueOrderSheetReceipt.Template(), constants.IssueOrderSheetReceipt.ArgsCount(), req.Args...)
		if err != nil {
			return r.Error(http.StatusInternalServerError, err.Error())
		}
		return r.Json(http.StatusOK, res)
	default:
		return r.Error(http.StatusBadRequest, "Bad Request : Not supported template")
	}
}

func validatePhoneNoInDev(phoneNo string) error {
	profile := os.Getenv("PROFILE")
	if profile == "dev" {
		VALID_PHONE_NUMBERS := []string{"01091751159", "01099457238", "01096782724", "01035192029", "01052493199"} // 김태완, 안중석, 김소진, 강수빈, 꿀벌원예 회장님

		count := 0
		for _, p := range VALID_PHONE_NUMBERS {
			if p == phoneNo {
				count += 1
			}
		}

		if count <= 0 {
			err := errors.New("invalid phone no to use in 'DEV' environment")
			return err
		} else {
			return nil
		}
	}
	return nil
}
