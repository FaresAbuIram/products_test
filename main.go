package main

import (
	"products/routes"
	"products/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	service.Connect()

	router := gin.Default()

	router.Use(cors.Default())
	routes.Setup(router)
	router.Run()
}
