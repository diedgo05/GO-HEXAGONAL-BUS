package routes

import (
	"bus-project/src/choferes/infraestructure/dependencies"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	routes := router.Group("/choferes")

	addChofer := dependencies.AddChoferController()
	getAllChofer := dependencies.GetAllChoferesController()
	updateChofer := dependencies.UpdateChoferController()
	deleteChofer := dependencies.DeleteChoferController()

	routes.POST("/", addChofer.Run)
	routes.GET("/", getAllChofer.Run)
	routes.PUT("/:id", updateChofer.Run)
	routes.DELETE("/:id", deleteChofer.Run)
}