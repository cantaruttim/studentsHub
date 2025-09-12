package main

import (
	"fmt"
	"time"
)

type Forms struct {
	nomeAluno      string
	matriculaAluno string
	emailAluno     string
	moduloAluno    string
	questionOne    string
	questionTwo    string
	setAt          time.Time
}

func (f Forms) GetFullForms() string {
	return fmt.Sprintf(
		"Name: %s\nregistrationNumber: %s\nEmail: %s\nModule: %s\nQuestion One: %s\nQuestion Two: %s\nCreated At: %s",
		f.nomeAluno,
		f.matriculaAluno,
		f.emailAluno,
		f.moduloAluno,
		f.questionOne,
		f.questionTwo,
		f.setAt.Format("2006-01-02 15:04:00"), // exibe só até os minutos
	)
}

func main() {
	now := time.Now().Truncate(time.Minute)

	form := Forms{
		nomeAluno:      "Matheus de Almeida Cantarutti",
		matriculaAluno: "2025024406",
		emailAluno:     "matheus.cantarutti@outlook.com",
		moduloAluno:    "Maturidade Espiritual",
		questionOne:    "Resposta da pergunta 1",
		questionTwo:    "Resposta da pergunta 2",
		setAt:          now,
	}

	fmt.Println(form.GetFullForms())
}
