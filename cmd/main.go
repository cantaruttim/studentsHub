package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type Forms struct {
	NomeAluno      string `json:"name"`
	MatriculaAluno string `json:"registrationNumber"`
	EmailAluno     string `json:"email"`
	ModuloAluno    string `json:"module"`
	QuestionOne    string `json:"questionOne"`
	QuestionTwo    string `json:"questionTwo"`
	SetAt          time.Time
}

func (f Forms) GetFullForms() string {
	return fmt.Sprintf(
		"Name: %s\nRegistrationNumber: %s\nEmail: %s\nModule: %s\nQuestion One: %s\nQuestion Two: %s\nCreated At: %s",
		f.NomeAluno,
		f.MatriculaAluno,
		f.EmailAluno,
		f.ModuloAluno,
		f.QuestionOne,
		f.QuestionTwo,
		f.SetAt.Format("2006-01-02 15:04:00"),
	)
}

func main() {
	server := gin.Default()
	server.GET("/forms", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.Run(":8000")

}
