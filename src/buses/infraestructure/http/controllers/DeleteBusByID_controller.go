package controllers

import (
	"bus-project/src/buses/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteBusByIDController struct {
	uc *application.DeleteBusByIDUseCase
}

func NewDeleteBusByIDController(uc *application.DeleteBusByIDUseCase) *DeleteBusByIDController {
	return &DeleteBusByIDController{uc: uc}
}

func (ctrl *DeleteBusByIDController) Run(c *gin.Context) {
	id := c.Param("idBus")
	busID, err := strconv.Atoi(id)	

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El id debe ser un número"})
		return
	}

	err = ctrl.uc.Run(busID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Bus eliminado correctamente"})
	}
}