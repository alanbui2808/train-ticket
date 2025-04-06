package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PongController struct{}

func NewPongController() *PongController {
	return &PongController{}
}

func (p *PongController) Pong(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{ // map of strings
		"message": "pong",
	})
}
