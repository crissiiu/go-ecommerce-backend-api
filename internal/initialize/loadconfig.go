package initialize

import (
	"fmt"

	"github.com/crisiiu/go-ecommerce-backend-api/global"
	"github.com/spf13/viper"
)

func LoadConfig() {
	viper := viper.New()
	viper.AddConfigPath("./configs") //path config
	viper.SetConfigName("local")     //tenfile
	viper.SetConfigType("yaml")

	//read file config
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fail to read configuration: %w", err))
	}

	//config structure
	if err := viper.Unmarshal(&global.Config); err != nil {
		panic(fmt.Errorf("Fail to unmarshal configuration: %w", err))
	}
}
