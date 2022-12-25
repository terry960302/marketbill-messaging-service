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

/** @{params}
- to : 받는 사람
- msg : 메세지 내용
- sendType : SMS, LMS, MMS 와 같은 메세지 포맷
*/
func (s *SmsService) SendDefaultSMS(to string, msg string, sendType string) (*models.SensResponse, error) {
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
			ErrLogs: errLog,
		}
		s.db.Create(&log)
	}()

	host := os.Getenv("SENS_HOST")
	serviceId := os.Getenv("SENS_SERVICE_ID")
	accessKeyId := os.Getenv("SENS_ACCESS_KEY_ID")
	supportedSendTypes := []string{constants.SMS, constants.LMS, constants.MMS}

	if !contains(supportedSendTypes, sendType) {
		msg := fmt.Sprintf("not supported send-type(%s). send-type should one of SMS, LMS, MMS", sendType)
		return nil, errors.New(msg)
	}

	path := "/sms/v2/services/" + serviceId + "/messages"
	url := host + path

	var reqBody models.SensRequest = models.SensRequest{
		Type:    sendType,
		From:    constants.FROM_PHONE_NO,
		Content: msg,
		Messages: []models.SensMessage{
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

	var smsResp models.SensResponse
	if err := json.Unmarshal(respBody, &smsResp); err != nil {
		return nil, err
	}

	return &smsResp, nil
}

// [to] : 받는 사람
// [template] : 메세지 템플릿
// [argsLength] : 템플릿에 들어가야하는 적정 args 개수
// [args] : 템플릿 내용에 필요한 값들
func (s *SmsService) SendSmsUsingTemplate(to string, template string, argsLength int, args ...interface{}) (*models.SensResponse, error) {
	if len(args) != argsLength {
		e := fmt.Sprintf("SendSmsUsingTemplate: Invalid args. There's must be %d args", argsLength)
		return nil, errors.New(e)
	}
	message := fmt.Sprintf(template, args...)
	var sendType string = constants.SMS

	bytes := []byte(message)
	if len(bytes) > constants.MAX_SMS_BYTES_LENGTH {
		sendType = constants.LMS
	}
	return s.SendDefaultSMS(to, message, sendType)
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

func contains(arr []string, element string) bool {
	for _, value := range arr {
		if value == element {
			return true
		}
	}
	return false
}
