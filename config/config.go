package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Config struct {
	HTTP struct {
		Port int
	}
	DB struct {
		DBURL string
	}
}

func LoadConfig(path string) (config Config, err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		} else {
			return config, errors.Wrap(err, "reading config file")
		}
	}

	viper.SetDefault("http.port", 8080)
	viper.SetDefault("db.dburl", "postgresql://postgres:postgres@postgres:5432/postgres")

	viper.AutomaticEnv()
	viper.SetEnvPrefix("STATS")
	viper.BindEnv("db.dburl", "DATABASE_URL")

	if err := viper.Unmarshal(&config); err != nil {
		return config, errors.Wrap(err, "unmarshaling config")
	}

	return config, nil
}
