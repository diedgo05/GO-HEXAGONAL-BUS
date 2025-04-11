package controllers

import (
	"bus-project/src/buses/application"
	"bus-project/src/buses/application/services"
	"bus-project/src/buses/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AddBusController struct {
	uc *application.AddBusUseCase
	event *services.Event
}

func NewAddBusController(uc *application.AddBusUseCase, event *services.Event) *AddBusController {
	return &AddBusController{uc: uc, event: event}
}

func (ctrl *AddBusController) Run(c *gin.Context) {
	var buses domain.Buses

	if err := c.ShouldBindJSON(&buses); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todos los campos son requeridos"})
		return
	}

	err := ctrl.uc.Run(buses)

	if err == nil {
		err = ctrl.event.Run(buses)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"status": true,
			"data": gin.H{
				"type": "buses",
				"idBus": buses.IdBus,
				"attributes": gin.H{
					"placa": buses.Placa,
					"capacidad": buses.Capacidad,
					"disponible": buses.Disponible,
					"chofer": buses.ChoferID,
				},
			},
		})
	}
}