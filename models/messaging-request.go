package models

type MessagingRequest struct {
	To          string        `json:"to"`
	MessageType string        `json:"message-type"`
	Args        []interface{} `json:"args"`
}
