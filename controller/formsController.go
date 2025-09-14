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

// fa = forms activities
func (fa *formsController) GetForms(ctx *gin.Context) {

	forms, err := fa.formsUsecase.GetForms()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, forms)
}

func (fa *formsController) FindById(ctx *gin.Context) {
	// Pega o valor da URL corretamente (case-sensitive)
	RN := ctx.Param("registrationNumber")

	if RN == "" {
		res := model.Respose{
			Message: "The value of RegistrationNumber must not be empty",
		}
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	forms, err := fa.formsUsecase.FindById(RN)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar no banco"})
		return
	}

	// Aqui valida se não encontrou nada
	if forms.RegistrationNumber == "" {
		res := model.Respose{
			Message: "The value of RegistrationNumber could not be found on database",
		}
		ctx.JSON(http.StatusNotFound, res)
		return
	}

	ctx.JSON(http.StatusOK, forms)
}

func (fa *formsController) CreateForms(ctx *gin.Context) {
	var form model.Forms

	err := ctx.BindJSON(&form)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	form.SentAt = time.Now()

	insertedForm, err := fa.formsUsecase.PostForms(form)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar no banco"})
		return
	}

	ctx.JSON(http.StatusCreated, insertedForm)
}
