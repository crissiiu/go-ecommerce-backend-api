package controller

import (
	"net/http"

	"github.com/crisiiu/go-ecommerce-backend-api/internal/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: services.NewUserService(),
	}
}

//controller -> Service -> Repo -> Model ->Db

func (uc *UserController) GetUserByID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": uc.userService.GetInfUser(),
	})
}
