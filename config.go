package main

import (
	"os"
	"os/user"
	"strconv"
	"syscall"

	"github.com/spf13/viper"
)

type Config struct {
	appKey     string
	appUrl     string
	configFile string
	disks      []string
	smartctl   string
}

func (config *Config) init() error {
	viper.SetConfigFile(config.configFile)
	viper.SetDefault("tools.smartctl", "smartctl")
	err := viper.ReadInConfig()
	if err != nil {
		return &errorString{"config: Got an error when trying to read config file - " + err.Error()}
	}
	config.appKey = viper.GetString("agent.key")
	config.appUrl = viper.GetString("agent.url")
	config.smartctl = viper.GetString("tools.smartctl")
	config.disks = viper.GetStringSlice("dev.disks")
	return nil
}

func NewConfig(configFile string) (*Config, error) {
	info, err := os.Stat(configFile)
	if err != nil {
		return nil, err
	}
	stat := info.Sys().(*syscall.Stat_t)
	user, err := user.Current()
	if err != nil {
		return nil, err
	}
	uid, err := strconv.Atoi(user.Uid)
	if err != nil {
		return nil, err
	}
	if stat.Uid != 0 && stat.Uid != uint32(uid) {
		return nil, &errorString{"The config file must be owned by current user or uid 0"}
	}
	if stat.Mode&3711 != 0 {
		return nil, &errorString{"The config file must have 0600 permissions"}
	}
	config := Config{configFile: configFile}
	err = config.init()
	if err != nil {
		return nil, err
	} else {
		return &config, nil
	}
}
