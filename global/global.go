package global

import (
	"github.com/alanbui/train-ticket/package/logger"
	"github.com/alanbui/train-ticket/package/setting"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config setting.Config    // stores all the config
	Logger *logger.LoggerZap // global logger
	Mdb    *gorm.DB
	Rdb    *redis.Client
)

/*
Config
Redis
Mysql

*/
