package controllers

import (
	"bus-project/src/choferes/application"
	"bus-project/src/choferes/domain"
	"net/http"
	"github.com/gin-gonic/gin"
)

type AddChoferController struct {
	uc *application.AddChoferUseCase
}

//Constructor de AddChoferController
func NewAddChoferController(uc *application.AddChoferUseCase) *AddChoferController {
	return &AddChoferController{uc: uc}
}

func (ctrl *AddChoferController) Run(c *gin.Context) {
	var choferes domain.Chofer

	if err := c.ShouldBindJSON(&choferes); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todos los campos son requeridos"})
		return
	} 

	err := ctrl.uc.Run(choferes)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"status": true,
			"data": gin.H{
				"type": "choferes",
				"idChofer": choferes.ID,
				"attributes": gin.H{
					"nombre": choferes.Nombre,
					"apellidos":choferes.Apellido_p + choferes.Apellido_m,
					"edad":choferes.Edad,
				},
			},
		})
	}
}