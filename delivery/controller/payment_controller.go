package controller

import (
	"livecode-wmb-2/delivery/api"
	"livecode-wmb-2/usecase"

	"github.com/gin-gonic/gin"
)

type PaymentController struct {
	router        *gin.Engine
	uclopei       usecase.LopeiUseCase
	ucTransaction usecase.TransactionUseCase
	ucTable       usecase.TableUseCase
	api.BaseApi
}

func NewPaymentController(router *gin.Engine, ucLopei usecase.LopeiUseCase, ucTransaction usecase.TransactionUseCase, ucTable usecase.TableUseCase) *PaymentController {
	// Disini akan terdapat kumpulan semua request method yang di butuhkan
	controller := PaymentController{
		router:        router,
		uclopei:       ucLopei,
		ucTransaction: ucTransaction,
		ucTable:       ucTable,
	}

	// ini method-methodnya
	// rgLopei := router.Group("api/lopei")
	// rgMenu.POST("/", controller.createNewMenu)
	// rgLopei.GET("/check-balance", controller.getBalance)
	// rgMenu.GET("/id", controller.getMenuById)
	// rgMenu.PUT("/update", controller.updateMenu)
	// rgMenu.DELETE("/delete", controller.deleteMenu)

	// rgMenuPrice := router.Group("api/menu-price")
	// rgMenuPrice.POST("/", controller.createAndUpdateNewMenuPrice)
	// rgMenuPrice.PUT("/", controller.createAndUpdateNewMenuPrice)

	return &controller
}
