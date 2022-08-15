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

type MenuController struct {
	router *gin.Engine
	ucmenu usecase.MenuUseCase
	api.BaseApi
}

func (p *MenuController) createNewMenu(c *gin.Context) {
	var newMenu model.Menu

	err := p.ParseRequestBody(c, &newMenu)
	if err != nil {
		p.Failed(c, utils.RequiredError())
		return
	}
	err = p.ucmenu.InsertMenu(&newMenu)
	if err != nil {
		p.Failed(c, err)
		return
	}
	p.Success(c, newMenu)

}

func (p *MenuController) createAndUpdateNewMenuPrice(c *gin.Context) {
	var newMenuPrice model.Menu
	id := c.Query("id")
	result, err := p.ucmenu.FindById(id)
	if err != nil {
		p.Failed(c, err)
		return
	}

	err = p.ParseRequestBody(c, &newMenuPrice)
	if err != nil {
		p.Failed(c, utils.RequiredError())
		return
	}

	result.MenuPrice = model.MenuPrice{
		Price: newMenuPrice.MenuPrice.Price,
	}

	err = p.ucmenu.UpdatePrice(&newMenuPrice)
	if err != nil {
		p.Failed(c, err)
		return
	}
	p.Success(c, newMenuPrice)

}

func (p *MenuController) getMenu(c *gin.Context) {

	res, err := p.ucmenu.ShowAllMenu("MenuPrice")
	if err != nil {
		p.Failed(c, err)
		return
	}
	p.Success(c, res)

}

func (p *MenuController) getMenuById(c *gin.Context) {
	id := c.Query("id")
	res, err := p.ucmenu.FindById(id)
	if err != nil {
		p.Failed(c, err)
		return
	}
	p.Success(c, res)

}

func (p *MenuController) updateMenu(c *gin.Context) {
	var existMenu model.Menu
	var newMenu map[string]interface{}

	id := c.Query("id")
	NewId := utils.ConverterStrToInt(id)
	existMenu.Model.ID = uint(NewId)

	err := p.ParseRequestBody(c, &newMenu)
	if err != nil {
		p.Failed(c, utils.RequiredError())
		return
	}
	err = p.ucmenu.UpdateMenu(
		&existMenu,
		newMenu)
	if err != nil {
		p.Failed(c, err)
		return
	}
	p.Success(c, newMenu)
}

func (p *MenuController) deleteMenu(c *gin.Context) {
	var existMenu model.Menu

	id := c.Query("id")
	NewId := utils.ConverterStrToInt(id)
	existMenu.Model.ID = uint(NewId)

	_, err := p.ucmenu.FindById(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		p.Failed(c, err)
		return
	}
	err = p.ucmenu.DeleteMenu(&existMenu)
	if err != nil {
		p.Failed(c, err)
		return
	}
	p.Success(c, nil)

}

func NewMenuController(router *gin.Engine, ucMenu usecase.MenuUseCase) *MenuController {
	// Disini akan terdapat kumpulan semua request method yang di butuhkan
	controller := MenuController{
		router: router,
		ucmenu: ucMenu,
	}

	// ini method-methodnya
	rgMenu := router.Group("api/menu", middleware.NewAuthTokenValidator(utils.NewTokenService(config.NewConfig().TokenConfig)).RequireToken())
	rgMenu.POST("/", controller.createNewMenu)
	rgMenu.GET("/", controller.getMenu)
	rgMenu.GET("/id", controller.getMenuById)
	rgMenu.PUT("/update", controller.updateMenu)
	rgMenu.DELETE("/delete", controller.deleteMenu)

	rgMenuPrice := router.Group("api/menu-price", middleware.NewAuthTokenValidator(utils.NewTokenService(config.NewConfig().TokenConfig)).RequireToken())
	rgMenuPrice.POST("/", controller.createAndUpdateNewMenuPrice)
	rgMenuPrice.PUT("/", controller.createAndUpdateNewMenuPrice)

	return &controller
}
