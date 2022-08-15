package controller

import (
	"fmt"
	"livecode-wmb-2/config"
	"livecode-wmb-2/delivery/api"
	"livecode-wmb-2/delivery/middleware"
	"livecode-wmb-2/model"
	"livecode-wmb-2/usecase"
	"livecode-wmb-2/utils"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	router        *gin.Engine
	ucTransaction usecase.TransactionUseCase
	ucTable       usecase.TableUseCase
	api.BaseApi
}

func (trx *TransactionController) orderTransaction(c *gin.Context) {
	var NewOrder model.Bill

	err := trx.ParseRequestBody(c, &NewOrder)
	if err != nil {
		trx.Failed(c, utils.RequiredError())
		return
	}

	status, err := trx.ucTable.FindById(NewOrder.TableID)
	if err != nil {
		trx.Failed(c, err)
		return
	}
	if !status.IsAvailable {
		fmt.Println("Maaf Table Sudah Terisi")
		return
	}
	err = trx.ucTransaction.CreateNewTransaction(&NewOrder)
	if err != nil {
		trx.Failed(c, err)
		return
	}
	existTable, err := trx.ucTable.FindById(NewOrder.TableID)
	if err != nil {
		trx.Failed(c, err)
		return
	}
	err = trx.ucTable.UpdateTable(
		&existTable,
		map[string]interface{}{
			"is_available": "False"})
	if err != nil {
		trx.Failed(c, err)
		return
	}
	trx.Success(c, NewOrder)
}

func (trx *TransactionController) paymentTransaction(c *gin.Context) {

	id := c.Query("id")
	newId := utils.ConverterStrToInt(id)
	NewPayment, err := trx.ucTransaction.PrintBillbyId(map[string]interface{}{"t_bill.id": newId})
	if err != nil {
		trx.Failed(c, err)
		return
	}
	res, err := trx.ucTransaction.FindById(newId)
	if err != nil {
		trx.Failed(c, err)
		return
	}
	existTable, err := trx.ucTable.FindById(res.TableID)
	if err != nil {
		trx.Failed(c, err)
		return
	}
	err = trx.ucTable.UpdateTable(
		&existTable,
		map[string]interface{}{
			"is_available": "True"})
	if err != nil {
		trx.Failed(c, err)
		return
	}
	trx.Success(c, NewPayment)
}

func (trx *TransactionController) getTransaction(c *gin.Context) {

	res, err := trx.ucTransaction.FindAll()
	if err != nil {
		trx.Failed(c, err)
		return
	}
	trx.Success(c, res)

}

func NewTransactionController(router *gin.Engine, ucTransaction usecase.TransactionUseCase, ucTable usecase.TableUseCase) *TransactionController {
	// Disini akan terdapat kumpulan semua request method yang di butuhkan
	controller := TransactionController{
		router:        router,
		ucTransaction: ucTransaction,
		ucTable:       ucTable,
	}

	// ini method-methodnya
	rgTransaction := router.Group("api/transaction", middleware.NewAuthTokenValidator(utils.NewTokenService(config.NewConfig().TokenConfig)).RequireToken())
	rgTransaction.POST("/order", controller.orderTransaction)
	rgTransaction.POST("/payment", controller.paymentTransaction)
	rgTransaction.GET("/", controller.getTransaction)
	// rgDiscount.GET("/id", controller.getDiscountById)
	// rgDiscount.PUT("/update", controller.updateDiscount)
	// rgDiscount.DELETE("/delete", controller.deleteDiscount)

	return &controller
}
