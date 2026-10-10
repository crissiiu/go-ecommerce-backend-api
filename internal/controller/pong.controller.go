package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PongController struct{}

func NewPong() *PongController {
	return &PongController{}
}

func (p *PongController) Pong(c *gin.Context) {
	// name := c.Param("name")
	name := c.DefaultQuery("name", "crissiiu")
	uid := c.Query("uid")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong" + name,
		"id":      uid,
		"users":   []string{"cr7", "m10", "crissiiu"},
	})
}
