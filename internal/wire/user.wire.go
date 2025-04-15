//go:build wireinject

package wire // warning should be ok for wire package

import (
	"github.com/alanbui/train-ticket/internal/controller"
	"github.com/alanbui/train-ticket/internal/repo"
	"github.com/alanbui/train-ticket/internal/service"
	"github.com/google/wire"
)

// used for Dependency Injection
func InitUserRouterHandler() (*controller.UserController, error) {
	wire.Build(
		repo.NewUserRepository,
		service.NewUserService,
		controller.NewUserController,
	)

	return new(controller.UserController), nil
}
