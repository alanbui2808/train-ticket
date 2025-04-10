package global

import (
	"github.com/alanbui/train-ticket/package/logger"
	"github.com/alanbui/train-ticket/package/setting"
)

var (
	Config setting.Config    // stores all the config
	Logger *logger.LoggerZap // global logger
)

/*
Config
Redis
Mysql

*/
