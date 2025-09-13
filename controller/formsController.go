package controller

import (
	"forms-activities/usecase"
	"net/http"

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

// fa = forms activities
func (fa *formsController) GetForms(ctx *gin.Context) {

	forms, err := fa.formsUsecase.GetForms()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	// forms := []model.Forms{
	// 	{
	// 		NomeAluno:      "Fernanda Oliveira da Silva",
	// 		MatriculaAluno: "2022003456",
	// 		EmailAluno:     "fernanda.silva@example.com",
	// 		ModuloAluno:    "Maturidade Espiritual",
	// 		QuestionOne:    "Como você descreveria sua jornada espiritual até agora?",
	// 		QuestionTwo:    "Quais práticas espirituais você gostaria de explorar mais?",
	// 		SentAt:         time.Now(),
	// 	},
	// }
	ctx.JSON(http.StatusOK, forms)
}
