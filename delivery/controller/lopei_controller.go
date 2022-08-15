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

type LopeiController struct {
	router  *gin.Engine
	uclopei usecase.LopeiUseCase
	api.BaseApi
}

func (l *LopeiController) getBalance(c *gin.Context) {
	var lopeiBalance model.Lopei
	id := c.Query("id")
	newId := utils.ConverterStrToInt(id)
	res, err := l.uclopei.GetBalance(int32(newId))
	// log.Println(res)
	if err != nil {
		l.Failed(c, err)
		return
	}
	lopeiBalance.Id = newId
	lopeiBalance.Balance = res
	l.Success(c, lopeiBalance)
}

func NewLopeiController(router *gin.Engine, ucLopei usecase.LopeiUseCase) *LopeiController {
	// Disini akan terdapat kumpulan semua request method yang di butuhkan
	controller := LopeiController{
		router:  router,
		uclopei: ucLopei,
	}

	rgLopei := router.Group("api/lopei", middleware.NewAuthTokenValidator(utils.NewTokenService(config.NewConfig().TokenConfig)).RequireToken())
	rgLopei.GET("/check-balance", controller.getBalance)

	return &controller
}
