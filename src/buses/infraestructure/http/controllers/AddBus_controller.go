package controllers

import (
	"bus-project/src/buses/application"
	"bus-project/src/buses/application/services"
	"bus-project/src/buses/domain"
	"net/http"
	"github.com/gin-gonic/gin"
)

type AddBusController struct {
	uc    *application.AddBusUseCase
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

	// Obtener el ID generado por el caso de uso
	id, err := ctrl.uc.Run(buses)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Actualizar el IdBus en el objeto buses
	buses.IdBus = id

	// Publicar el evento con el objeto actualizado
	err = ctrl.event.Run(buses)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{buses.Placa: buses})
}