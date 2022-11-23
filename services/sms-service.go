package services

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"marketbill-messaging-service/constants"
	"marketbill-messaging-service/models"
	"marketbill-messaging-service/utils"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func SendDefaultSMS(to string, msg string) (*models.SmsResponse, error) {
	// host := config.C.Api.Host
	// serviceId := config.C.Api.ServiceId
	host := os.Getenv("SENS_HOST")
	serviceId := os.Getenv("SENS_SERVICE_ID")
	accessKeyId := os.Getenv("SENS_ACCESS_KEY_ID")

	path := "/sms/v2/services/" + serviceId + "/messages"
	url := host + path

	var reqBody models.SmsRequest = models.SmsRequest{
		Type:    "SMS",
		From:    constants.FROM_PHONE_NO,
		Content: msg,
		Messages: []models.SmsMessage{
			{
				To: to,
			},
		},
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}
	bodyReader := bytes.NewReader(jsonBody)

	req, err := http.NewRequest("POST", url, bodyReader)
	if err != nil {
		return nil, err
	}

	defer req.Body.Close()

	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	signature := generateSignature("POST", path, timestamp, accessKeyId)

	req.Header.Add("Content-type", "application/json")
	req.Header.Add("x-ncp-apigw-timestamp", strconv.Itoa(int(timestamp)))
	req.Header.Add("x-ncp-iam-access-key", accessKeyId)
	req.Header.Add("x-ncp-apigw-signature-v2", signature)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var smsResp models.SmsResponse
	if err := json.Unmarshal(respBody, &smsResp); err != nil {
		return nil, err
	}

	return &smsResp, nil
}

func generateSignature(method string, path string, timestamp int64, accessKey string) string {
	secretKey := os.Getenv("SENS_SECRET_KEY")
	// secretKey := config.C.Api.SecretKey
	bodyList := []string{
		method,
		" ",
		path,
		"\n",
		strconv.Itoa(int(timestamp)),
		"\n",
		accessKey,
	}
	body := strings.Join(bodyList, "")
	sig := utils.HMAC256(body, secretKey)
	return sig
}
