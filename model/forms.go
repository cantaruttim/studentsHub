package model

import (
	"time"
)

type Forms struct {
	MatriculaAluno string    `json:"registrationNumber"`
	NomeAluno      string    `json:"name"`
	EmailAluno     string    `json:"email"`
	ModuloAluno    string    `json:"module"`
	QuestionOne    string    `json:"questionOne"`
	QuestionTwo    string    `json:"questionTwo"`
	SentAt         time.Time `json:"createdAt"`
}
