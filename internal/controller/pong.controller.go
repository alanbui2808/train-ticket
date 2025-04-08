package controller

import (
	"fmt"

	"github.com/alanbui/train-ticket/package/response"
	"github.com/gin-gonic/gin"
)

type PongController struct{}

func NewPongController() *PongController {
	return &PongController{}
}
func (p *PongController) Pong(ctx *gin.Context) {
	fmt.Println("===> My Handler")
	response.SuccessResponse(ctx, 20001, nil)
}
