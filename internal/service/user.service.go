package service

import (
	"github.com/alanbui/train-ticket/internal/repo"
	"github.com/alanbui/train-ticket/package/response"
)

// UserRepo embedded inside UserService
// type UserService struct {
// 	userRepo *repo.UserRepo
// }

// // Service => Repo
// func NewUserService() *UserService {
// 	return &UserService{
// 		userRepo: repo.NewUserRepo(),
// 	}
// }

// // user service us
// func (us *UserService) GetUserInfoService() string {
// 	return us.userRepo.GetUserInfo()
// }

// INTERFACE

type IUserService interface {
	Register(email string, purpose string) int
}

// now private (everything works through UserService interface)
type userService struct {
	userRepo repo.IUserRepository
}

func NewUserService(userRepo repo.IUserRepository) IUserService {
	return &userService{userRepo: userRepo}
}

// Register implements IUserService.
// turn userService into IUserService
func (us *userService) Register(email string, purpose string) int {
	if us.userRepo.GetUserByEmail(email) {
		return response.ErrCodeUserExists
	}
	return response.ErrCodeSuccess
}
