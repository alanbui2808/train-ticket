package routers

// // route the requests to correct APIs
// import (
// 	"fmt"

// 	c "github.com/alanbui/train-ticket/internal/controller"
// 	"github.com/alanbui/train-ticket/internal/middlewares"
// 	"github.com/gin-gonic/gin"
// )

// func AA() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		fmt.Println("Before ==> AA")
// 		c.Next() // continue to next Handler
// 		fmt.Println("Alfter ==> AA")
// 	}
// }

// func BB() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		fmt.Println("Before ==> BB")
// 		c.Next() // continue to next Handler
// 		fmt.Println("Alfter ==> BB")
// 	}
// }

// func CC(c *gin.Context) {
// 	fmt.Println("Before ==> CC")
// 	c.Next() // continue to next Handler
// 	fmt.Println("Alfter ==> CC")
// }

// // must have capital letter so that main can acess.
// func NewRouter() *gin.Engine {
// 	r := gin.Default()
// 	// use the middleware
// 	/*
// 		Before AA
// 		Before BB
// 		Before CC
// 		--> Handler!
// 		After CC
// 		After BB
// 		After AA
// 	*/
// 	r.Use(middlewares.AuthenMiddleware(), BB(), CC)
// 	v1 := r.Group("v1/2025")
// 	{
// 		// v1/2025/user/1
// 		v1.GET("ping", c.NewPongController().Pong) // v1/2025/ping
// 		v1.GET("/user/1", c.NewUserController().GetUserbyID)
// 	}

// 	return r
// }
