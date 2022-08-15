package delivery

import (
	"livecode-wmb-2/config"
	"livecode-wmb-2/delivery/controller"
	"livecode-wmb-2/manager"

	"github.com/gin-gonic/gin"
)

type appServer struct {
	usecaseManager manager.UseCaseManager
	// tokenService   utils.Token
	engine *gin.Engine
	host   string
}

func Server() *appServer {
	r := gin.Default()
	appConfig := config.NewConfig()
	infra := manager.NewInfra(appConfig)
	repoManager := manager.NewRepositoryManager(infra)
	useCaseManager := manager.NewUseCaseManager(repoManager)
	// tokenService := utils.NewTokenService(appConfig.TokenConfig)

	host := appConfig.Url
	return &appServer{
		usecaseManager: useCaseManager,
		// tokenService:   tokenService,
		engine: r,
		host:   host,
	}
}

func (a *appServer) initControllers() {
	controller.NewMenuController(a.engine, a.usecaseManager.MenuUseCase())
	controller.NewTableController(a.engine, a.usecaseManager.TableUseCase())
	controller.NewTransTypeController(a.engine, a.usecaseManager.TransType())
	controller.NewDiscountController(a.engine, a.usecaseManager.Discount())
	controller.NewCustomerController(a.engine, a.usecaseManager.Customer())
	controller.NewTransactionController(a.engine, a.usecaseManager.Transaction(), a.usecaseManager.TableUseCase())
	controller.NewLopeiController(a.engine, a.usecaseManager.LopeiChekBalance())
	controller.NewLoginController(a.engine)
}

func (a *appServer) Run() {

	a.initControllers()
	err := a.engine.Run(a.host)
	if err != nil {
		panic(err)
	}
}
