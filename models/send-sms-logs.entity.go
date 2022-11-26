package models

import "gorm.io/gorm"

type SendSmsLogs struct {
	gorm.Model
	To      string `json:"to"`
	Message string `json:"message"`
	Status  string `json:"status"` // SUCCESS, FAILURE
	ErrLogs string `gorm:"type:text" json:"err_logs"`
}
