package main

import (
	dependenciesBus "bus-project/src/buses/infraestructure/dependenciesBus"
	routesBus "bus-project/src/buses/infraestructure/http/routesBus"
	choferDependencies "bus-project/src/choferes/infraestructure/dependencies"
	choferRoutes "bus-project/src/choferes/infraestructure/http/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"}, // Origen de tu frontend Angular
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Métodos permitidos
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"}, // Encabezados permitidos
		ExposeHeaders:    []string{"Content-Length"}, // Encabezados expuestos
		AllowCredentials: true, // Permitir credenciales (cookies, auth headers, etc.)
	}))

    // Inicializar dependencias de choferes
    choferDependencies.Init()
    choferRoutes.SetupRoutes(r)

    // Inicializar dependencias de buses
	dependenciesBus.InitBus()
    routesBus.Routes(r)

    r.Run()
}