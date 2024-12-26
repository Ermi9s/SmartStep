package main

import (
	"log"
	"os"
	"step-project/database"
	"step-project/delivery/controller"
	"step-project/delivery/router"
	"step-project/usecase"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Panic("Failed to load .env" , err.Error())
	}

	port := os.Getenv("PORT");

	if port == "" {
		port = "8080"
	}

	var service database.ServerConnection
	service.Connect_could()
	client := service.Client
	devicecollection := client.Database("StepDB").Collection("Devices")

	usc := usecase.New_Usecase(devicecollection)
	ctrl := controller.New_Controller(usc)

	Router := gin.Default()
	router.SetupRoutes(Router , ctrl)
	Router.Run("0.0.0.0:"+port)
}