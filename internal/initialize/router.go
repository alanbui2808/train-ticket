package initialize

// route the requests to correct APIs
import (
	"github.com/alanbui/train-ticket/global"
	"github.com/alanbui/train-ticket/internal/routers"
	"github.com/gin-gonic/gin"
)

// must have capital letter so that main can acess.
func InitRouter() *gin.Engine {
	var r *gin.Engine

	// set the respective server mode
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	// middlewares
	managerRouter := routers.RouterGroupApp.Manager
	userRouter := routers.RouterGroupApp.User

	MainGroup := r.Group("v1/2025")
	{
		MainGroup.GET("/checkStatus") // monitor status
	}
	// set groupings for user router
	{
		userRouter.InitUserRouter(MainGroup)
		userRouter.InitProductRouter(MainGroup)
	}
	// set groupings for manager router
	{
		managerRouter.InitUserRouter(MainGroup)
		managerRouter.InitAdminRouter(MainGroup)
	}

	return r

}
