package router

import (
	"backend-engineer-test-privy/cake/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter(cakeController controller.CakeController) *gin.Engine {
	r := gin.Default()
	r.Use(Middleware())

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "API service is ready!")
	})

	cake := r.Group("/cake")
	{
		cake.POST("", cakeController.CreateCake)
		cake.GET("", cakeController.GetAllCakes)
		cake.GET("/:id", cakeController.GetCakeByID)
		cake.PATCH("/:id", cakeController.UpdateCake)
		cake.DELETE("/:id", cakeController.DeleteCake)
	}

	return r
}
