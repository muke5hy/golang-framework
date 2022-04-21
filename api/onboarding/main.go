package main

import (
	"onboarding/configs"
	"onboarding/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//run database
	configs.ConnectDB()
	//routes
	routes.UserRoute(router) //add this

	router.Run("localhost:6000")
}
