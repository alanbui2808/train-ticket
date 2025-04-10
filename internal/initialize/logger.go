package initialize

import (
	"github.com/alanbui/train-ticket/global"
	"github.com/alanbui/train-ticket/package/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
