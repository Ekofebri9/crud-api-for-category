package configs

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Port string `mapstructure:"PORT"`
}

func Init() *Config {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	config := Config{
		Port: viper.GetString("PORT"),
	}
	return &config
}
