package controller

import (
	"livecode-wmb-2/config"
	"livecode-wmb-2/delivery/api"
	"livecode-wmb-2/model"
	"livecode-wmb-2/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	router *gin.Engine
	api.BaseApi
}

func (lc *LoginController) login(c *gin.Context) {
	tokenService := utils.NewTokenService(config.NewConfig().TokenConfig)
	var user model.Credential
	if err := c.BindJSON(&user); err != nil {
		lc.Failed(c, err)
		return
	}
	if user.Username == "admin@example.com" && user.Password == "12345678" {
		token, err := tokenService.CreateAccesToken(&user)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		lc.Success(c, token)
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}

func NewLoginController(router *gin.Engine) *LoginController {
	// Disini akan terdapat kumpulan semua request method yang di butuhkan
	controller := LoginController{
		router: router,
	}

	// ini method-methodnya
	rgLogin := router.Group("api/auth")
	rgLogin.POST("/login", controller.login)
	// rgCustomer.GET("/", controller.getCustomer)
	// rgCustomer.PUT("/activation", controller.activationCustomer)
	// rgDiscount.PUT("/update", controller.updateDiscount)
	// rgDiscount.DELETE("/delete", controller.deleteDiscount)

	return &controller
}
