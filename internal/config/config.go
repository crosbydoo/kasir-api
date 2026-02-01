package config

import (
	"kasir-api/internal/pkg"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Port   string
	DBConn string
}

func LoadConfig() *Config {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		if err := viper.ReadInConfig(); err != nil {
			pkg.Log.WithFields(logrus.Fields{
				"error": err.Error(),
			}).Warn("Failed to read .env file, using environment variables")
		}
	}

	port := viper.GetString("PORT")
	if port == "" {
		port = "8080"
	}

	config := &Config{
		Port:   port,
		DBConn: viper.GetString("DB_CONN"),
	}

	return config
}
