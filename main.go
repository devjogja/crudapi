package main

import (
	"log"

	routes "github.com/devjogja/crudapi/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.Router()
	log.Fatal(router.run(":4747"))
}
