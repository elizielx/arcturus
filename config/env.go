package config

import "github.com/spf13/viper"

type Configuration struct {
	DatabaseHost     string `mapstructure:"POSTGRES_HOST"`
	DatabasePort     string `mapstructure:"POSTGRES_PORT"`
	DatabaseUser     string `mapstructure:"POSTGRES_USER"`
	DatabasePassword string `mapstructure:"POSTGRES_PASSWORD"`
	DatabaseName     string `mapstructure:"POSTGRES_DB"`
	ServerPort       string `mapstructure:"SERVER_PORT"`
}

func LoadConfiguration(path string) (Configuration, error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName(".env")
	viper.AutomaticEnv()

	var configuration Configuration
	err := viper.ReadInConfig()
	if err != nil {
		return Configuration{}, err
	}

	err = viper.Unmarshal(&configuration)
	if err != nil {
		return Configuration{}, err
	}

	return configuration, nil
}
