package main

import (
	"log"

	"github.com/devjogja/crudapi/config"
	routes "github.com/devjogja/crudapi/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Connect DB
	config.Connect()
	// Init Router
	router := gin.Default()
	// Route Handlers / Endpoints
	routes.Routes(router)
	log.Fatal(router.Run(":4747"))
}
