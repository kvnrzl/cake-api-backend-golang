//go:build wireinject
// +build wireinject

package main

import (
	"backend-engineer-test-privy/cake/controller"
	"backend-engineer-test-privy/cake/repository/mysql"
	"backend-engineer-test-privy/cake/service"
	"backend-engineer-test-privy/database"
	"backend-engineer-test-privy/router"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
)

func InitServer() *gin.Engine {
	wire.Build(
		database.InitDBMysql,
		validator.New,

		mysql.NewCakeRepositoryImpl,
		service.NewCakeServiceImpl,
		controller.NewCakeControllerImpl,

		router.SetupRouter,
	)

	return nil
}
