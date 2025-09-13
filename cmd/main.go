package main

import (
	"forms-activities/controller"
	"forms-activities/db"
	"forms-activities/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	FormsUsecase := usecase.NewFormsUsecase()

	// Controllers
	formsController := controller.NewFormsController(FormsUsecase)

	server.GET("/api/v1/forms-activities/response", formsController.GetForms)

	server.Run(":8000")

}
