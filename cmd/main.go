package main

import (
	"forms-activities/controller"
	"forms-activities/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	FormsUsecase := usecase.NewFormsUsecase()

	// Controllers
	formsController := controller.NewFormsController(FormsUsecase)

	server.GET("/api/v1/forms-activities/response", formsController.GetForms)

	server.Run(":8000")

}
