package models

type DefaultSmsRequest struct {
	// From    string `json:"from"`
	To      string `json:"to"`
	Message string `json:"message"`
}
