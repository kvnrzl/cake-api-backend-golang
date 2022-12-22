package controller

import (
	"backend-engineer-test-privy/cake/service"
	"backend-engineer-test-privy/helper"
	"backend-engineer-test-privy/model"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type CakeControllerImpl struct {
	cakeService service.CakeService
}

func NewCakeControllerImpl(cakeService service.CakeService) CakeController {
	return &CakeControllerImpl{cakeService: cakeService}
}

func (a *CakeControllerImpl) CreateCake(c *gin.Context) {
	cake := &model.Cake{}

	err := c.ShouldBindJSON(cake)
	if err != nil {
		logrus.Error(err)
		helper.ResponseBadRequest(c, err.Error())
		return
	}

	cake, err = a.cakeService.CreateCake(c, cake)
	if errors.Is(err, model.ErrInputFieldInvalid) {
		logrus.Error(err)
		helper.ResponseUnprocessableEntity(c, err.Error())
		return
	}

	if errors.Is(err, model.ErrCannotCreate) {
		logrus.Error(err)
		helper.ResponseInternalServerError(c, err.Error())
		return
	}

	if err != nil {
		logrus.Error(err)
		helper.ResponseInternalServerError(c, err.Error())
		return
	}

	helper.ResponseStatusCreated(c, cake)
}

func (a *CakeControllerImpl) GetCakeByID(c *gin.Context) {
	id := c.Param("id")
	parseID, err := strconv.Atoi(id)
	if err != nil {
		logrus.Error(err)
		helper.ResponseBadRequest(c, "Cannot parse the ID that you provide")
		return
	}

	uintID := uint(parseID)
	cake, err := a.cakeService.GetCakeByID(c, uintID)
	if errors.Is(err, model.ErrRecordNotFound) {
		logrus.Error(err)
		helper.ResponseCakeNotFound(c, uintID)
		return
	}

	if err != nil {
		logrus.Error(err)
		helper.ResponseInternalServerError(c, err.Error())
		return
	}

	helper.ResponseStatusOK(c, cake)
}

func (a *CakeControllerImpl) GetAllCakes(c *gin.Context) {
	cakes, err := a.cakeService.GetAllCakes(c)
	if err != nil {
		logrus.Error(err)
		helper.ResponseInternalServerError(c, err.Error())
		return
	}

	helper.ResponseStatusOK(c, cakes)
}

func (a *CakeControllerImpl) UpdateCake(c *gin.Context) {
	id := c.Param("id")
	parseID, err := strconv.Atoi(id)
	if err != nil {
		logrus.Error(err)
		helper.ResponseBadRequest(c, "Cannot parse the ID that you provide")
		return
	}

	cake := &model.Cake{}

	err = c.ShouldBindJSON(cake)
	if err != nil {
		logrus.Error(err)
		helper.ResponseBadRequest(c, err.Error())
		return
	}

	uintID := uint(parseID)
	cake.ID = uintID
	cake, err = a.cakeService.UpdateCake(c, cake)
	if errors.Is(err, model.ErrInputFieldInvalid) {
		logrus.Error(err)
		helper.ResponseUnprocessableEntity(c, err.Error())
		return
	}

	if errors.Is(err, model.ErrRecordNotFound) {
		logrus.Error(err)
		helper.ResponseCakeNotFound(c, uintID)
		return
	}

	if errors.Is(err, model.ErrCannotUpdate) {
		logrus.Error(err)
		helper.ResponseInternalServerError(c, err.Error())
		return
	}

	if err != nil {
		logrus.Error(err)
		helper.ResponseInternalServerError(c, err.Error())
		return
	}

	helper.ResponseStatusOK(c, cake)
}

func (a *CakeControllerImpl) DeleteCake(c *gin.Context) {
	id := c.Param("id")
	parseID, err := strconv.Atoi(id)
	if err != nil {
		logrus.Error(err)
		helper.ResponseBadRequest(c, "Cannot parse the ID that you provide")
		return
	}

	uintID := uint(parseID)
	err = a.cakeService.DeleteCake(c, uintID)
	if errors.Is(err, model.ErrRecordNotFound) {
		logrus.Error(err)
		helper.ResponseCakeNotFound(c, uintID)
		return
	}

	if errors.Is(err, model.ErrCannotDelete) {
		logrus.Error(err)
		helper.ResponseInternalServerError(c, err.Error())
		return
	}

	if err != nil {
		logrus.Error(err)
		helper.ResponseInternalServerError(c, err.Error())
		return
	}

	helper.ResponseStatusOK(c, struct{}{})
}
