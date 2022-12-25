package models

type MessagingRequest struct {
	To       string        `json:"to"`
	Template string        `json:"template"`
	Args     []interface{} `json:"args"`
}
