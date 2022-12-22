package controller

import "github.com/gin-gonic/gin"

type CakeController interface {
	CreateCake(c *gin.Context)
	GetCakeByID(c *gin.Context)
	GetAllCakes(c *gin.Context)
	UpdateCake(c *gin.Context)
	DeleteCake(c *gin.Context)
}
