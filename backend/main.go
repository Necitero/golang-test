package main

import (
	"github.com/gin-gonic/gin"
	"necitero/golang-test/routes"
)

func main() {
	router := gin.Default()
	routes.RegisterRoutes(router)
	router.Run()
}
