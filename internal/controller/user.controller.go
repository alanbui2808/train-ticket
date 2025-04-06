package controller

import (
	"net/http"

	"github.com/alanbui/train-ticket/internal/service"
	"github.com/gin-gonic/gin"
)

// Controller => Service => Repo
// Repo => Models => DB
type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

func (uc *UserController) GetUserbyID(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{ // map of strings
		"message": uc.userService.GetUserInfoService(),
		"users":   []string{"Faker", "Chovy"},
	})
}
