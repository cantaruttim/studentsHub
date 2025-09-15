package controller

import (
	"forms-activities/dto"
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

func (fa *formsController) ReceiveForms(ctx *gin.Context) {
	var form dto.FormsDTO

	err := ctx.BindJSON(&form)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	form.SentAtDTO = time.Now()

	modelForm := form.ToModel()

	insertedForm, err := fa.formsUsecase.PostForms(modelForm)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar no banco"})
		return
	}

	ctx.JSON(http.StatusCreated, insertedForm)
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
	RN := ctx.Param("registrationNumber")

	if RN == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "registrationNumber must not be empty"})
		return
	}

	form, err := fa.formsUsecase.FindById(RN)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar no banco"})
		return
	}

	if form == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Formulário não encontrado"})
		return
	}

	ctx.JSON(http.StatusOK, form)
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
