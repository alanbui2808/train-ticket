package global

import (
	"github.com/alanbui/train-ticket/package/logger"
	"github.com/alanbui/train-ticket/package/setting"
	"gorm.io/gorm"
)

var (
	Config setting.Config    // stores all the config
	Logger *logger.LoggerZap // global logger
	Mdb    *gorm.DB
)

/*
Config
Redis
Mysql

*/
