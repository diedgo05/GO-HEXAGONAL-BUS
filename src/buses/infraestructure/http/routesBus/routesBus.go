package routes

import (
	"bus-project/src/buses/infraestructure/dependenciesBus"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	routes := router.Group(("/buses"))

	addBus := dependencies.AddBusController()
	getAllBus := dependencies.GetAllBusesController()
	updateBus := dependencies.UpdateBusController()
	getBusByIdChofer := dependencies.GetBusByIdChoferController()
	deleteBus := dependencies.DeleteBusController()

	routes.POST("/", addBus.Run)
	routes.GET("/", getAllBus.Run)
	routes.PUT("/:idBus", updateBus.Run)
	routes.GET("/:choferID", getBusByIdChofer.Run)
	routes.DELETE("/:idBus", deleteBus.Run)
}