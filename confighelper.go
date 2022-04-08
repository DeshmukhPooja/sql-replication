package confighelper

import (
	"fmt"

	"github.com/spf13/viper"
)

func InitViper() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
}
