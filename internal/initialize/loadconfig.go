package initialize

import (
	"fmt"

	"github.com/alanbui/train-ticket/global"
	"github.com/spf13/viper"
)

func LoadConfig() {
	viper := viper.New()
	viper.AddConfigPath("./config/") // path to config
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	// read configuration
	err := viper.ReadInConfig()
	if err != nil {
		// quit the program immediately
		panic(fmt.Errorf("failed to read config %v", err))
	}

	// viper reads directly from yaml file.
	// fmt.Println("Server port::", viper.GetInt("server.port"))
	// fmt.Println("Security key::", viper.GetString("security.jwt.key"))

	// configure structure
	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("Unable to decode configuration %v", err)
	}
}
