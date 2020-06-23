package config

import (
	"log"

	"github.com/spf13/viper"
)

func init(){
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("Fatal error config file: ", err)
	}
}