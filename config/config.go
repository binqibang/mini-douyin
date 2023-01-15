package config

import (
	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	App      *App      `yaml:"app"`
	Database *Database `yaml:"database"`
}

type App struct {
	Address string `yaml:"address"`
}

type Database struct {
	Mysql *Mysql `yaml:"mysql"`
}

type Mysql struct {
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	Address  string `yaml:"address"`
	Database string `yaml:"database"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(configFile string) (*Config, error) {
	var config *Config
	viper.SetConfigType("yaml")
	viper.SetConfigFile(configFile)

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	} else {
		return config, nil
	}
}
