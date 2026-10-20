package initialize

import (
	"github.com/crisiiu/go-ecommerce-backend-api/global"
	"github.com/crisiiu/go-ecommerce-backend-api/internal/routers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	var router *gin.Engine

	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		router = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		router = gin.New()
	}

	// middelwares
	router.Use() //Logging
	router.Use() //Cross
	router.Use() // Litmit Global

	manageRouter := routers.RouterGroupApp.Manage
	userRouter := routers.RouterGroupApp.User

	MainGroup := router.Group("/v1/2026")
	{
		MainGroup.GET("/check-status")
	}
	{
		userRouter.InitUserRouter(MainGroup)
		userRouter.InitProductRouter(MainGroup)
	}
	{
		manageRouter.InitUserRouter(MainGroup)
		manageRouter.InitAdminRouter(MainGroup)
	}
	return router
}
