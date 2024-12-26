package domain

import "github.com/gin-gonic/gin"

type Main_Controller_Interface interface {
	GetAllInfo(c *gin.Context)
	RegisterDevice(c *gin.Context)
}

type Main_Usecase_Interface interface {
	GetAllInfo() ([]LocationData, error)
	RegisterDevice(Device) error
}