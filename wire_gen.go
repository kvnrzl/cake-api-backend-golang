// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"backend-engineer-test-privy/cake/controller"
	"backend-engineer-test-privy/cake/repository/mysql"
	"backend-engineer-test-privy/cake/service"
	"backend-engineer-test-privy/database"
	"backend-engineer-test-privy/router"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Injectors from injector.go:

func InitServer() *gin.Engine {
	cakeRepository := mysql.NewCakeRepositoryImpl()
	db := database.InitDBMysql()
	validate := validator.New()
	cakeService := service.NewCakeServiceImpl(cakeRepository, db, validate)
	cakeController := controller.NewCakeControllerImpl(cakeService)
	engine := router.SetupRouter(cakeController)
	return engine
}