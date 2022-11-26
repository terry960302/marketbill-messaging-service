package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"marketbill-messaging-service/constants"
	"marketbill-messaging-service/models"
	"marketbill-messaging-service/utils"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
)

type SmsService struct {
	db *gorm.DB
}

func NewSmsService(db *gorm.DB) *SmsService {
	return &SmsService{db: db}
}

func (s *SmsService) SendDefaultSMS(to string, msg string) (*models.SmsResponse, error) {
	defer func() {
		var status string = constants.SUCCESS
		var errLog string = ""
		var err error = nil

		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New("unknown panic")
			}
			status = constants.FAILURE
			errLog = err.Error()
		}

		log := models.SendSmsLogs{
			To:      to,
			Message: msg,
			Status:  status,
			Log:     errLog,
		}
		s.db.Save(log)
	}()

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

func (s *SmsService) SendSmsUsingTemplate(to string, template string, args ...interface{}) (*models.SmsResponse, error) {
	if len(args) != 1 {
		return nil, errors.New("SendVerificationSms: Invalid args. There's must be 1 args")
	}
	message := fmt.Sprintf(template, args...)
	return s.SendDefaultSMS(to, message)
}

func generateSignature(method string, path string, timestamp int64, accessKey string) string {
	secretKey := os.Getenv("SENS_SECRET_KEY")
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
