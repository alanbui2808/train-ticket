package initialize

import (
	"github.com/alanbui/train-ticket/global"
	"github.com/alanbui/train-ticket/package/logger"
)

func InitLogger() {
	// initialize global Logger to a custom Logger that we defined.
	global.Logger = logger.NewLogger(global.Config.Logger)
}
