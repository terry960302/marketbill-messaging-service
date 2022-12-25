package models

type DefaultMessagingRequest struct {
	// From    string `json:"from"`
	To      string `json:"to"`
	Message string `json:"message"`
}
