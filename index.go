package main

import (
	"./config"
)

func main() {
	router := config.Router()
	router.Run(":3000")
}
