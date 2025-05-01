package initialize

import (
	"fmt"

	"github.com/alanbui/train-ticket/global"
	"go.uber.org/zap"
)

// contains all the initializers
func Run() {
	// load config
	LoadConfig()
	m := global.Config.Mysql
	fmt.Println("Loading config mysql", m.Username, m.Password)
	InitLogger()
	global.Logger.Info("Config Log ok!!", zap.String("ok", "success"))
	InitMySql()
	InitRedis()

	r := InitRouter()

	r.Run() // 8080
}
