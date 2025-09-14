package main

import (
	"forms-activities/controller"
	"forms-activities/db"
	"forms-activities/repository"
	"forms-activities/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	// Repository
	FormsRepository := repository.NewFormsRepository(dbConnection)

	// Usecase
	FormsUsecase := usecase.NewFormsUsecase(FormsRepository)

	// Controllers
	formsController := controller.NewFormsController(FormsUsecase)

	server.GET("/api/v1/forms-activities/response", formsController.GetForms)
	server.POST("/api/v1/forms-activities/response", formsController.CreateForms)

	server.Run(":8000")

}
