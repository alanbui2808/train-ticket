package routers

// route the requests to correct APIs
import (
	c "github.com/alanbui/train-ticket/internal/controller"
	"github.com/gin-gonic/gin"
)

// must have capital letter so that main can acess.
func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("v1/2025")
	{
		// v1/2025/user/1
		v1.GET("/user/1", c.NewUserController().GetUserbyID)
	}

	return r
}
