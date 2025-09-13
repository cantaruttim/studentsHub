package main

import (
	"forms-activities/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	// Controllers
	formsController := controller.NewFormsController()

	server.GET("/api/v1/forms-activities/response", formsController.GetForms)

	server.Run(":8000")

}
