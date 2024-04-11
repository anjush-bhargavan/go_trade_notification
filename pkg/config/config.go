package config

import "github.com/spf13/viper"

// Config represents the environment variables.
type Config struct {
	Email    string `mapstructure:"EMAIL"`
	Password string `mapstructure:"PASSWORD"`
}

// LoadConfig will load the environment variable to access.
func LoadConfig() *Config {
	var config Config
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	viper.Unmarshal(&config)
	return &config
}
