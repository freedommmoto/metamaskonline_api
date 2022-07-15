package tool

import (
	"fmt"
	"github.com/spf13/viper"
)

var configValue ConfigObject

//all config from .env
type ConfigObject struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOUECE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	LINEToken     string `mapstructure:"LINEToken"`
	RegisterUrl   string `mapstructure:"REGISTER_URL"`
}
type makeNewCustomer struct {
	CustomerName string `json:"customer_name" binding:"required"`
	Currency     string `json:"currency" binding:"required,oneof=USD EUR" `
}

func LoadConfig(part string) (config ConfigObject, err error) {
	viper.AddConfigPath(part)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		fmt.Printf("unable to decode into struct %v", err)
	}
	configValue = config
	return
}
