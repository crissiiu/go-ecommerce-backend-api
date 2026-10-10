package controller

import (
	"github.com/crisiiu/go-ecommerce-backend-api/internal/services"
	"github.com/crisiiu/go-ecommerce-backend-api/pkg/response"
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
	response.SuccessResponse(c, 20001, "Thanh Cong")
}
