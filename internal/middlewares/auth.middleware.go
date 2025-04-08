package middlewares

import (
	"github.com/alanbui/train-ticket/package/response"
	"github.com/gin-gonic/gin"
)

func AuthenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "valid-token" {
			response.ErrorResponse(c, response.ErrInvalidToken, "")
			c.Abort() // disrupt the chain
			return
		}
		c.Next()
	}
}
