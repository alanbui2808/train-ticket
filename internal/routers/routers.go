package routers

// route the requests to correct APIs
import (
	c "github.com/alanbui/train-ticket/internal/controller"
	"github.com/gin-gonic/gin"
)

// must have capital letter so that main can acess.
func NewRouter() *gin.Engine {
	r := gin.Default()

	// In this way we know Pong belongs to NewPongController.
	// This is why PongController struct{} is useful.
	r.GET("/ping", c.NewPongController().Pong)

	return r
}
