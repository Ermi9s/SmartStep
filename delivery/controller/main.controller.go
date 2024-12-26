package controller

import (
	"net/http"
	"step-project/domain"
	"github.com/gin-gonic/gin"
)

type Main_Controller struct {
	Usecase domain.Main_Usecase_Interface
}

func New_Controller(usecase domain.Main_Usecase_Interface) *Main_Controller {
	return &Main_Controller{
		Usecase: usecase,
	}
}

func (mc *Main_Controller)GetAllInfo(c *gin.Context) {
	data, err := mc.Usecase.GetAllInfo()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

func (mc *Main_Controller)RegisterDevice(c *gin.Context) {
	var device domain.Device

	if err := c.BindJSON(&device); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := mc.Usecase.RegisterDevice(device)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Device registered successfully"})
}