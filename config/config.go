package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	DbConfig DbConfig `mapstructure:"database"`
}

type DbConfig struct {
	Name     string
	Host     string
	Port     string
	Username string
	Password string
}

func Load(configName, configType, configPath string) Config {
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	viper.AddConfigPath(configPath)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("failed to read config : %v", err)
		os.Exit(1)
	}

	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		fmt.Printf("failed to unmarshal config : %v", err)
		os.Exit(1)
	}

	return c
}
