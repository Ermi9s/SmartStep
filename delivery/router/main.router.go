package router

import (
	"step-project/delivery/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") 
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func SetupRoutes(router *gin.Engine, mainController *controller.Main_Controller) {
	router.Use(CORSMiddleware())

	api := router.Group("/api")
	{
		api.GET("/info", mainController.GetAllInfo)
		api.POST("/register", mainController.RegisterDevice)
		api.GET("/update", mainController.UpdateStatus)
	}
}
