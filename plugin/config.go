package plugin

import (
	"github.com/spf13/viper"
)

type Config struct {
	*viper.Viper
}

func NewConfig() *Config {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	v.AddConfigPath("./config")

	// Set defaults
	v.SetDefault("api_port", "8080")
	v.SetDefault("grpc_port", "50051")

	if err := v.ReadInConfig(); err != nil {
		// Use defaults if config file not found
	}

	return &Config{v}
}
