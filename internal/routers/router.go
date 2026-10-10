package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/v1/2026")
	v1.GET("/ping", Pong) //v1/2026/ping
	v1.POST("/ping", Pong)
	v1.PUT("/ping", Pong)
	v1.PATCH("/ping", Pong)
	v1.DELETE("/ping", Pong)
	v1.OPTIONS("/ping", Pong)

	return router
}

func Pong(c *gin.Context) {
	// name := c.Param("name")
	name := c.DefaultQuery("name", "crissiiu")
	uid := c.Query("uid")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong" + name,
		"id":      uid,
		"users":   []string{"cr7", "m10", "crissiiu"},
	})
}
