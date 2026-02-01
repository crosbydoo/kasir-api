package pkg

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func InitLogger() {
	Log = logrus.New()

	// Set output ke stdout
	Log.SetOutput(os.Stdout)

	// Set format JSON untuk production, atau Text untuk development
	env := os.Getenv("APP_ENV")
	if env == "production" {
		Log.SetFormatter(&logrus.JSONFormatter{})
	} else {
		Log.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05",
			ForceColors:     true,
		})
	}

	// Set log level
	Log.SetLevel(logrus.InfoLevel)

	Log.Info("Logger initialized successfully")
}
