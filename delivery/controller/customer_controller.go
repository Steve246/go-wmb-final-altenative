package controller

import (
	"livecode-wmb-2/config"
	"livecode-wmb-2/delivery/api"
	"livecode-wmb-2/delivery/middleware"
	"livecode-wmb-2/model"
	"livecode-wmb-2/usecase"
	"livecode-wmb-2/utils"

	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	router     *gin.Engine
	ucCustomer usecase.CustomerUseCase
	api.BaseApi
}

func (cs *CustomerController) createNewCustomer(c *gin.Context) {
	var NewCustomer model.Customer

	err := cs.ParseRequestBody(c, &NewCustomer)
	if err != nil {
		cs.Failed(c, utils.RequiredError())
		return
	}
	err = cs.ucCustomer.Registration(&NewCustomer)
	if err != nil {
		cs.Failed(c, err)
		return
	}
	cs.Success(c, NewCustomer)
}

func (cs *CustomerController) getCustomer(c *gin.Context) {

	res, err := cs.ucCustomer.GetAllCustomer("Discount")
	if err != nil {
		cs.Failed(c, err)
		return
	}
	cs.Success(c, res)

}

func (cs *CustomerController) activationCustomer(c *gin.Context) {
	var NewDiscount model.Discount
	err := cs.ParseRequestBody(c, &NewDiscount)
	if err != nil {
		cs.Failed(c, utils.RequiredError())
		return
	}
	id := c.Query("id")
	newId := utils.ConverterStrToInt(id)
	err = cs.ucCustomer.ActivationMemberForExistingCustomerAndAddDiscount(newId, NewDiscount.Description, NewDiscount.Pct)
	if err != nil {
		cs.Failed(c, err)
		return
	}
}

func NewCustomerController(router *gin.Engine, ucCustomer usecase.CustomerUseCase) *CustomerController {
	// Disini akan terdapat kumpulan semua request method yang di butuhkan
	controller := CustomerController{
		router:     router,
		ucCustomer: ucCustomer,
	}

	// ini method-methodnya
	rgCustomer := router.Group("api/customer", middleware.NewAuthTokenValidator(utils.NewTokenService(config.NewConfig().TokenConfig)).RequireToken())
	rgCustomer.POST("/registration", controller.createNewCustomer)
	rgCustomer.GET("/", controller.getCustomer)
	rgCustomer.PUT("/activation", controller.activationCustomer)
	// rgDiscount.PUT("/update", controller.updateDiscount)
	// rgDiscount.DELETE("/delete", controller.deleteDiscount)

	return &controller
}
