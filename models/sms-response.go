package models

type SmsResponse struct {
	RequestId   string `json:"requestId"`
	RequestTime string `json:"requestTime"`
	StatusCode  string `json:"statusCode"`
	StatusName  string `json:"statusName"`
}
