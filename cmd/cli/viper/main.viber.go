package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Databases []struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
	} `mapstructure:"databases"`
}

func main() {
	viper := viper.New()
	viper.AddConfigPath("./configs") //path config
	viper.SetConfigName("local")     //tenfile
	viper.SetConfigType("yaml")

	//read file config
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fail to read configuration: %w", err))
	}

	//read server configuration
	fmt.Println("Sever port: ", viper.GetInt("server.port"))
	fmt.Println("JWT Key: ", viper.GetString("security.jwt.key"))

	//config structure
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		panic(fmt.Errorf("Fail to unmarshal configuration: %w", err))
	}

	fmt.Println("Server port from struct: ", config.Server.Port)
	for i, db := range config.Databases {
		fmt.Printf("Database %d:\n", i+1)
		fmt.Printf("  User: %s\n", db.User)
		fmt.Printf("  Password: %s\n", db.Password)
		fmt.Printf("  Host: %s\n", db.Host)
	}
}
