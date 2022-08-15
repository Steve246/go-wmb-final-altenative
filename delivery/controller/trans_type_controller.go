package controller

import (
	"errors"
	"livecode-wmb-2/config"
	"livecode-wmb-2/delivery/api"
	"livecode-wmb-2/delivery/middleware"
	"livecode-wmb-2/model"
	"livecode-wmb-2/usecase"
	"livecode-wmb-2/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TransTypeController struct {
	router      *gin.Engine
	ucTransType usecase.TransTypeUseCase
	api.BaseApi
}

func (t *TransTypeController) createNewTransType(c *gin.Context) {
	var NewTransType model.TransType

	err := t.ParseRequestBody(c, &NewTransType)
	if err != nil {
		t.Failed(c, utils.RequiredError())
		return
	}
	err = t.ucTransType.CreateNewTransType(&NewTransType)
	if err != nil {
		t.Failed(c, err)
		return
	}
	t.Success(c, NewTransType)
}

func (t *TransTypeController) getTransType(c *gin.Context) {

	res, err := t.ucTransType.FindAllTransType()
	if err != nil {
		t.Failed(c, err)
		return
	}
	t.Success(c, res)

}

func (t *TransTypeController) getTransTypeById(c *gin.Context) {
	id := c.Query("id")
	newId := utils.ConverterStrToInt(id)
	res, err := t.ucTransType.FindById(newId)
	if err != nil {
		t.Failed(c, err)
		return
	}
	t.Success(c, res)

}

func (t *TransTypeController) updateTransType(c *gin.Context) {
	var existTransType model.TransType
	var newTransType map[string]interface{}

	id := c.Query("id")
	NewId := utils.ConverterStrToInt(id)
	existTransType.Model.ID = uint(NewId)

	err := t.ParseRequestBody(c, &newTransType)
	if err != nil {
		t.Failed(c, utils.RequiredError())
		return
	}
	err = t.ucTransType.UpdateTransType(
		&existTransType,
		newTransType)
	if err != nil {
		t.Failed(c, err)
		return
	}
	t.Success(c, newTransType)
}

func (t *TransTypeController) deleteTransType(c *gin.Context) {
	var existTransType model.TransType

	id := c.Query("id")
	NewId := utils.ConverterStrToInt(id)
	existTransType.Model.ID = uint(NewId)

	_, err := t.ucTransType.FindById(NewId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		t.Failed(c, err)
		return
	}
	err = t.ucTransType.DeleteTransType(&existTransType)
	if err != nil {
		t.Failed(c, err)
		return
	}
	t.Success(c, nil)

}

func NewTransTypeController(router *gin.Engine, ucTransType usecase.TransTypeUseCase) *TransTypeController {
	// Disini akan terdapat kumpulan semua request method yang di butuhkan
	controller := TransTypeController{
		router:      router,
		ucTransType: ucTransType,
	}

	// ini method-methodnya
	rgTransType := router.Group("api/trans-type", middleware.NewAuthTokenValidator(utils.NewTokenService(config.NewConfig().TokenConfig)).RequireToken())
	rgTransType.POST("/", controller.createNewTransType)
	rgTransType.GET("/", controller.getTransType)
	rgTransType.GET("/id", controller.getTransTypeById)
	rgTransType.PUT("/update", controller.updateTransType)
	rgTransType.DELETE("/delete", controller.deleteTransType)

	return &controller
}
