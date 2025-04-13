package user

import "github.com/gin-gonic/gin"

type ProductRouter struct{}

func (pr *ProductRouter) InitProductRouter(Router *gin.RouterGroup) {
	// public routers
	productRouterPublic := Router.Group("/product")
	{
		productRouterPublic.GET("/search")
		productRouterPublic.GET("/details/:id")
	}

	// private routers

}
