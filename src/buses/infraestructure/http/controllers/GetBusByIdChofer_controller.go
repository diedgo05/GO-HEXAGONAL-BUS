package controllers

import (
	"bus-project/src/buses/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetBusByIdChoferController struct {
	uc *application.FindBusByIdChoferUseCase
}

func NewGetBusByIdChoferController(uc *application.FindBusByIdChoferUseCase) *GetBusByIdChoferController {
	return &GetBusByIdChoferController{uc: uc}
}

func (ctrl *GetBusByIdChoferController) Run(c *gin.Context) {
	id := c.Param("choferID")
	ChoferID, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El id debe ser un número"})
		return 
	}

	buses, err := ctrl.uc.Run(ChoferID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"buses": buses})
	}
}