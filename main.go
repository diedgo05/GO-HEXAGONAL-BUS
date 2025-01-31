package main

import (
	dependenciesBus "bus-project/src/buses/infraestructure/dependenciesBus"
	routesBus "bus-project/src/buses/infraestructure/http/routesBus"
	choferDependencies "bus-project/src/choferes/infraestructure/dependencies"
	choferRoutes "bus-project/src/choferes/infraestructure/http/routes"

	"github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    // Inicializar dependencias de choferes
    choferDependencies.Init()
    choferRoutes.SetupRoutes(r)

    // Inicializar dependencias de buses
	dependenciesBus.InitBus()
    routesBus.Routes(r)

    r.Run()
}