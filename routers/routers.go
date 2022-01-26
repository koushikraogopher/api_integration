package routers

import (
	"api_integration/client"
	"api_integration/controllers"

	"github.com/gin-gonic/gin"
)

// NewRouter will handle all routes
func NewRouter() *gin.Engine {

	// set server mode
	// gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	client := client.NewClient()
	getController := controllers.NewPetsController(client)

	v1 := router.Group("/v1")
	{
		v1.GET("/getPets", getController.GetPets)
		v1.GET("/getPets/:id", getController.GetPetById)
		v1.POST("/postPets", getController.PostPets)

	}

	return router
}
