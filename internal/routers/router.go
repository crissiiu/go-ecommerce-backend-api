package routers

import (
	"github.com/crisiiu/go-ecommerce-backend-api/internal/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/v1/2026")
	v1.GET("/ping", controller.NewPong().Pong)                  //v1/2026/ping
	v1.GET("/user", controller.NewUserController().GetUserByID) //v1/2026/ping
	// v1.POST("/ping", Pong)
	// v1.PUT("/ping", Pong)
	// v1.PATCH("/ping", Pong)
	// v1.DELETE("/ping", Pong)
	// v1.OPTIONS("/ping", Pong)

	return router
}
