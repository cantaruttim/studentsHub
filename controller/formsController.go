package controller

import (
	"forms-activities/model"
	"forms-activities/usecase"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type formsController struct {
	formsUsecase usecase.FormsUsecase
}

func NewFormsController(formsUsecase usecase.FormsUsecase) formsController {
	return formsController{
		formsUsecase: formsUsecase,
	}
}

func (p *formsController) GetForms(ctx *gin.Context) {
	forms := []model.Forms{
		{
			NomeAluno:      "Fernanda Oliveira da Silva",
			MatriculaAluno: "2022003456",
			EmailAluno:     "fernanda.silva@example.com",
			ModuloAluno:    "Maturidade Espiritual",
			QuestionOne:    "Como você descreveria sua jornada espiritual até agora?",
			QuestionTwo:    "Quais práticas espirituais você gostaria de explorar mais?",
			SentAt:         time.Now(),
		},
	}
	ctx.JSON(http.StatusOK, forms)
}
