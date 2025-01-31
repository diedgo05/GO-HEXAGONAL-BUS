package controllers

import (
	"bus-project/src/buses/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllBusesController struct {
	uc *application.GetAllBusesUseCase
}

func NewGetAllBusesController(uc *application.GetAllBusesUseCase) *GetAllBusesController {
	return &GetAllBusesController{uc: uc}
}

func (ctrl *GetAllBusesController) Run(c *gin.Context) {
	buses, err := ctrl.uc.Run()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error":"Todos los campos son requeridos",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"buses": buses})
}