package routes

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jekhedhri/tuto006/application/ui/controllers"
)

//Init is init
func Init() {
	router := gin.Default()
	v1 := router.Group("/api/v1/product")
	{
		v1.GET("/find/id=:id", controllers.FindProduct)
		v1.GET("/create", controllers.CreateProduct)

	}
	port := os.Getenv("PORT")
	router.Run(":" + port)
}
