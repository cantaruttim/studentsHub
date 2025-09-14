package model

import (
	"time"
)

type Forms struct {
	RegistrationNumber string    `json:"RegistrationNumber"`
	Name               string    `json:"Name"`
	Email              string    `json:"Email"`
	Module             string    `json:"Module"`
	QuestionOne        string    `json:"QuestionOne"`
	QuestionTwo        string    `json:"QuestionTwo"`
	SentAt             time.Time `json:"-"`
}
