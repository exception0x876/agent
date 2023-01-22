package main

import (
	"github.com/spf13/viper"
)

type Config struct {
	appKey     string
	appUrl     string
	configFile string
	disks      []string
}

func (config *Config) init() error {
	viper.SetConfigFile(config.configFile)
	err := viper.ReadInConfig()
	if err != nil {
		return &errorString{"config: Got an error when trying to read config file - " + err.Error()}
	}
	config.appKey = viper.GetString("agent.key")
	config.appUrl = viper.GetString("agent.url")
	config.disks = viper.GetStringSlice("dev.disks")
	return nil
}

func NewConfig(configFile string) (*Config, error) {
	config := Config{configFile: configFile}
	err := config.init()
	if err != nil {
		return nil, err
	} else {
		return &config, nil
	}
}
