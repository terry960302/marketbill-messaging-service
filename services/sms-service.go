package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"marketbill-messaging-service/config"
	"marketbill-messaging-service/models"
	"net/http"
)

func SendDefaultSMS(from string, to string, msg string) error {
	host := config.C.Api.Host
	serviceId := config.C.Api.ServiceId
	url := host + "/sms/v2/services/" + serviceId + "/messages"
	contentType := "application/json"

	var req models.SmsRequest = models.SmsRequest{
		Type:    "SMS",
		From:    from,
		Content: msg,
		Messages: []models.SmsMessage{
			{
				To: to,
			},
		},
	}

	jsonBody, err := json.Marshal(req)
	if err != nil {
		return err
	}
	resp, err := http.Post(url, contentType, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(respBody))

	var smsResp models.SmsResponse
	if err := json.Unmarshal(respBody, &smsResp); err != nil {
		return err
	}

	fmt.Println(smsResp)
	return nil
}
