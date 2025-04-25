package manager

import "github.com/gin-gonic/gin"

type UserRouter struct{}

/*
manager/user.router.go
Purpose: admin managing user (activate/deactivate, ban, reset_password user)
Access: Only by authenticated admins (behind middlewares) - no public routes and avoid authen
*/
func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// public routers
	// userRouterPublic := Router.Group("/admin/user")
	// {
	// 	userRouterPublic.GET("/register")
	// 	userRouterPublic.GET("/otp")
	// }

	// private routers (requires login)
	userRouterPrivate := Router.Group("/admin/user")
	// userRouterPrivate.Use(limiter())
	// userRouterPrivate.Use(Authen())
	// userRouterPrivate.Use(Permission())
	{
		userRouterPrivate.POST("/active_user_") // manage users
	}
}
