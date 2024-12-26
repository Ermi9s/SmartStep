package router

import (
	"step-project/delivery/controller"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, mainController *controller.Main_Controller) {
	api := router.Group("/api")
	{
		api.GET("/info", mainController.GetAllInfo)   
		api.POST("/register", mainController.RegisterDevice) 
	}
}