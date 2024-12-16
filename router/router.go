package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Creates a gin router with default middleware:
	router := gin.Default()

	// Define the routes
	initializeRoutes(router)

	// Define the routes
	router.Run() // listen and serve on
}
