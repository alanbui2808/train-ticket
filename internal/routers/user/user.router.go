package user

import (
	"github.com/alanbui/train-ticket/internal/wire"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

/*
user/user.router.go
Purpose: Actual end-user self-service routes.
For users activing on themselves (register, get info, etc)

Access: Can have public routes (e.g onboarding or verification) and some private required login
*/
func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// dependency injection
	userController, _ := wire.InitUserRouterHandler()

	// public routers
	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.POST("/register", userController.Register)
		userRouterPublic.POST("/otp")
	}

	// private routers (requires login)
	userRouterPrivate := Router.Group("/user")
	// userRouterPrivate.Use(limiter())
	// userRouterPrivate.Use(Authen())
	// userRouterPrivate.Use(Permission())
	{
		userRouterPrivate.GET("/get_info")
	}
}
