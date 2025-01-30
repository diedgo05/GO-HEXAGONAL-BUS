package main

import (
	"bus-project/src/choferes/infraestructure/dependencies"
	"bus-project/src/choferes/infraestructure/http/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	dependencies.Init()
	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run()
}