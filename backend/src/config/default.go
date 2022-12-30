// Package config handles the configuration
package config

import (
	"github.com/spf13/viper"
)

// Config struct holds the config values read
type Config struct {
	PostgreDriver  string `mapstructure:"POSTGRES_DRIVER"`
	PostgresSource string `mapstructure:"POSTGRES_SOURCE"`

	MongoURL      string `mapstructure:"MONGO_URL"`
	MongoUsername string `mapstructure:"MONGO_INITDB_ROOT_USERNAME"`
	MongoPassword string `mapstructure:"MONGO_INITDB_ROOT_PASSWORD"`

	Port string `mapstructure:"PORT"`
}

// LoadConfig loads .env file from the given path
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName(".env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	viper.SetEnvPrefix("openfridge")
	viper.BindEnv("POSTGRES_DRIVER")

	err = viper.Unmarshal(&config)
	return
}
