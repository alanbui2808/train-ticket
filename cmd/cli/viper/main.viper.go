package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	// anynomous inline struct declaration
	// this struct doesnt have a name. Note: Server is a field name in Config
	Server struct {
		Port int `mapstructure:"port"`
		// tell Viper to map to server in yaml file
	} `mapstructure:"server"` // Viper maps to server
	Databases []struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
	} `mapstructure:"databases"` // map to database
}

func main() {
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
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Unable to decode configuration %v", err)
	}

	// fmt.Println("Config Port::", config.Server.Port)
	// fmt.Print(config.Databases)
	for _, db := range config.Databases {
		fmt.Printf("databases User: %s, pass: %s, host: %s \n", db.User, db.Password, db.Host)
	}
}
