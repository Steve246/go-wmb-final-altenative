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

type DiscountController struct {
	router     *gin.Engine
	ucDiscount usecase.DiscountUseCase
	api.BaseApi
}

func (d *DiscountController) createNewDiscount(c *gin.Context) {
	var NewDiscount model.Discount

	err := d.ParseRequestBody(c, &NewDiscount)
	if err != nil {
		d.Failed(c, utils.RequiredError())
		return
	}
	err = d.ucDiscount.CreateNewDiscount(&NewDiscount)
	if err != nil {
		d.Failed(c, err)
		return
	}
	d.Success(c, NewDiscount)
}

func (d *DiscountController) getDiscount(c *gin.Context) {

	res, err := d.ucDiscount.FindAllDiscount()
	if err != nil {
		d.Failed(c, err)
		return
	}
	d.Success(c, res)

}

func (d *DiscountController) getDiscountById(c *gin.Context) {
	id := c.Query("id")
	newId := utils.ConverterStrToInt(id)
	res, err := d.ucDiscount.FindById(newId)
	if err != nil {
		d.Failed(c, err)
		return
	}
	d.Success(c, res)

}

func (d *DiscountController) updateDiscount(c *gin.Context) {
	var existDiscount model.Discount
	var newDiscount map[string]interface{}

	id := c.Query("id")
	NewId := utils.ConverterStrToInt(id)
	existDiscount.Model.ID = uint(NewId)

	err := d.ParseRequestBody(c, &newDiscount)
	if err != nil {
		d.Failed(c, utils.RequiredError())
		return
	}
	err = d.ucDiscount.UpdateDiscount(
		&existDiscount,
		newDiscount)
	if err != nil {
		d.Failed(c, err)
		return
	}
	d.Success(c, newDiscount)
}

func (d *DiscountController) deleteDiscount(c *gin.Context) {
	var existDiscount model.Discount

	id := c.Query("id")
	NewId := utils.ConverterStrToInt(id)
	existDiscount.Model.ID = uint(NewId)

	_, err := d.ucDiscount.FindById(NewId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		d.Failed(c, err)
		return
	}
	err = d.ucDiscount.DeleteDiscount(&existDiscount)
	if err != nil {
		d.Failed(c, err)
		return
	}
	d.Success(c, nil)

}

func NewDiscountController(router *gin.Engine, ucDiscount usecase.DiscountUseCase) *DiscountController {
	// Disini akan terdapat kumpulan semua request method yang di butuhkan
	controller := DiscountController{
		router:     router,
		ucDiscount: ucDiscount,
	}

	// ini method-methodnya
	rgDiscount := router.Group("api/discount", middleware.NewAuthTokenValidator(utils.NewTokenService(config.NewConfig().TokenConfig)).RequireToken())
	rgDiscount.POST("/", controller.createNewDiscount)
	rgDiscount.GET("/", controller.getDiscount)
	rgDiscount.GET("/id", controller.getDiscountById)
	rgDiscount.PUT("/update", controller.updateDiscount)
	rgDiscount.DELETE("/delete", controller.deleteDiscount)

	return &controller
}
