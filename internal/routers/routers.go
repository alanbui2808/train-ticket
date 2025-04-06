package routers

// route the requests to correct APIs
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// must have capital letter so that main can acess.
func NewRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", Pong)

	return r
}

// API
func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{ // map of strings
		"message": "pong",
	})
}
