package controllers

import (
	"bus-project/src/choferes/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetChoferByIdController struct {
	uc *application.GetChoferByIDUseCase
}

func NewGetChoferByIDController(uc *application.GetChoferByIDUseCase) *GetChoferByIdController{
	return &GetChoferByIdController{uc:uc}
}

func (ctrl *GetChoferByIdController) Run(c *gin.Context) {
	id := c.Param("id")
	ID, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El id debe ser un número"})
		return
	}

	chofer, err := ctrl.uc.Run(ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"chofer encontrado por id": chofer})
	}
}