package main

import (
	"github.com/gin-gonic/gin"
	"necitero/golang-test/database"
	"necitero/golang-test/routes"
)

func main() {
	database.SetupDatabase()
	router := gin.Default()
	routes.RegisterRoutes(router)
	router.Run()
}
