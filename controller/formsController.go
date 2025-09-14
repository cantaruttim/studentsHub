package controller

import (
	"forms-activities/model"
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
	ctx.JSON(http.StatusOK, forms)
}

func (fa *formsController) CreateForms(ctx *gin.Context) {
	var form model.Forms
	err := ctx.BindJSON(&form)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedForm, err := fa.formsUsecase.PostForms(form)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedForm)

}
