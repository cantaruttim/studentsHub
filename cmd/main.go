package main

import (
	"forms-activities/controller"
	"forms-activities/db"
	"forms-activities/repository"
	"forms-activities/usecase"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	// Habilita CORS padrão (aceita tudo)
	server.Use(cors.Default())

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
	server.GET("/api/v1/forms-activities/response/:registrationNumber", formsController.FindById)
	// server.POST("/api/v1/forms-activities/response", formsController.CreateForms)
	server.POST("/api/v1/forms-activities/response", formsController.ReceiveForms)

	server.Run(":8000")
}
