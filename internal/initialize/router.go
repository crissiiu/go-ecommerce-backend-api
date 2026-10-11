package initialize

import (
	"fmt"

	"github.com/crisiiu/go-ecommerce-backend-api/internal/controller"
	"github.com/crisiiu/go-ecommerce-backend-api/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func AA() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before --> AA")
		c.Next()
		fmt.Println("After --> AA")
	}
}

func BB() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before --> BB")
		c.Next()
		fmt.Println("After --> BB")
	}
}

func CC(c *gin.Context) {
	fmt.Println("Before --> CC")
	c.Next()
	fmt.Println("After --> CC")
}

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.Use(middlewares.AuthenticationMiddleware(), BB(), CC)
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
