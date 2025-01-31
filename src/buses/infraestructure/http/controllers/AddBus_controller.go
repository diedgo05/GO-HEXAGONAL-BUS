package controllers

import (
	"bus-project/src/buses/application"
	"bus-project/src/buses/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AddBusController struct {
	uc *application.AddBusUseCase
}

func NewAddBusController(uc *application.AddBusUseCase) *AddBusController {
	return &AddBusController{uc: uc}
}

func (ctrl *AddBusController) Run(c *gin.Context) {
	var buses domain.Buses

	if err := c.ShouldBindJSON(&buses); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todos los campos son requeridos"})
		return
	}

	err := ctrl.uc.Run(buses)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"status": true,
			"data": gin.H{
				"type": "Bus",
				"idBus": buses.IdBus,
				"attributes": gin.H{
					"placa": buses.Placa,
					"capacidad": buses.Capacidad,
					"chofer": buses.ChoferID,
				},
			},
		})
	}
}