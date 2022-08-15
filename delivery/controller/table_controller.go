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

type TableController struct {
	router  *gin.Engine
	ucTable usecase.TableUseCase
	api.BaseApi
}

func (t *TableController) createNewTable(c *gin.Context) {
	var NewTable model.Table

	err := t.ParseRequestBody(c, &NewTable)
	if err != nil {
		t.Failed(c, utils.RequiredError())
		return
	}
	err = t.ucTable.CreateNewTable(&NewTable)
	if err != nil {
		t.Failed(c, err)
		return
	}
	t.Success(c, NewTable)
}

func (t *TableController) getTable(c *gin.Context) {

	res, err := t.ucTable.FindAll()
	if err != nil {
		t.Failed(c, err)
		return
	}
	t.Success(c, res)

}

func (t *TableController) getTableById(c *gin.Context) {
	id := c.Query("id")
	newId := utils.ConverterStrToInt(id)
	res, err := t.ucTable.FindById(newId)
	if err != nil {
		t.Failed(c, err)
		return
	}
	t.Success(c, res)

}

func (t *TableController) updateTable(c *gin.Context) {
	var existTable model.Table
	var newTable map[string]interface{}

	id := c.Query("id")
	NewId := utils.ConverterStrToInt(id)
	existTable.Model.ID = uint(NewId)

	err := t.ParseRequestBody(c, &newTable)
	if err != nil {
		t.Failed(c, utils.RequiredError())
		return
	}
	err = t.ucTable.UpdateTable(
		&existTable,
		newTable)
	if err != nil {
		t.Failed(c, err)
		return
	}
	t.Success(c, newTable)
}

func (t *TableController) deleteTable(c *gin.Context) {
	var existTable model.Table

	id := c.Query("id")
	NewId := utils.ConverterStrToInt(id)
	existTable.Model.ID = uint(NewId)

	_, err := t.ucTable.FindById(NewId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		t.Failed(c, err)
		return
	}
	err = t.ucTable.DeleteTable(&existTable)
	if err != nil {
		t.Failed(c, err)
		return
	}
	t.Success(c, nil)

}

func NewTableController(router *gin.Engine, ucTable usecase.TableUseCase) *TableController {
	// Disini akan terdapat kumpulan semua request method yang di butuhkan
	controller := TableController{
		router:  router,
		ucTable: ucTable,
	}

	// ini method-methodnya
	rgTable := router.Group("api/table", middleware.NewAuthTokenValidator(utils.NewTokenService(config.NewConfig().TokenConfig)).RequireToken())
	rgTable.POST("/", controller.createNewTable)
	rgTable.GET("/", controller.getTable)
	rgTable.GET("/id", controller.getTableById)
	rgTable.PUT("/update", controller.updateTable)
	rgTable.DELETE("/delete", controller.deleteTable)

	return &controller
}
