package model

import (
	"time"
)

type Forms struct {
	id             int       `json:"id"`
	NomeAluno      string    `json:"name"`
	MatriculaAluno string    `json:"registrationNumber"`
	EmailAluno     string    `json:"email"`
	ModuloAluno    string    `json:"module"`
	QuestionOne    string    `json:"questionOne"`
	QuestionTwo    string    `json:"questionTwo"`
	SentAt         time.Time `json:"createdAt"`
}
