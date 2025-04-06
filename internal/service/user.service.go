package service

import "github.com/alanbui/train-ticket/internal/repo"

// UserRepo embedded inside UserService
type UserService struct {
	userRepo *repo.UserRepo
}

// Service => Repo
func NewUserService() *UserService {
	return &UserService{
		userRepo: repo.NewUserRepo(),
	}
}

// user service us
func (us *UserService) GetUserInfoService() string {
	return us.userRepo.GetUserInfo()
}
