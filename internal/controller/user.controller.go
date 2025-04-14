package controller

import (
	"github.com/alanbui/train-ticket/internal/service"
	"github.com/alanbui/train-ticket/package/response"
	"github.com/gin-gonic/gin"
)

// type UserController struct {
// 	userService *service.UserService
// }

// func NewUserController() *UserController {
// 	return &UserController{
// 		userService: service.NewUserService(),
// 	}
// }

// // Controller => Service => Repo
// // Repo => Models => DB

// func (uc *UserController) GetUserbyID(ctx *gin.Context) {
// 	response.SuccessResponse(ctx, 20001, []string{"Faker", "Chovy"})
// }

// INTERFACE
type UserController struct {
	userService service.IUserService
}

func NewUserController(userService service.IUserService) *UserController {
	return &UserController{userService: userService}
}

func (uc *UserController) Register(c *gin.Context) {
	result := uc.userService.Register("", "")
	response.SuccessResponse(c, result, nil)
}
