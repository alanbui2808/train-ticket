package initialize

// contains all the initializers
func Run() {
	// load config
	LoadConfig()
	InitLogger()
	InitMySql()
	InitRedis()

	r := InitRouter()

	r.Run()
}
