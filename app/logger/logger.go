package logger

import (
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Init() *logrus.Logger {
	var logger = logrus.New()
	logger.Formatter = &logrus.JSONFormatter{}

	if !viper.GetBool("APP_DEBUG") {
		logger.Out = ioutil.Discard
	}

	switch viper.GetString("LOG_CHANNEL") {
	case "file":
		// create directory
		var path = viper.GetString("LOG_PATH")
		if _, err := os.Stat(path); os.IsNotExist(err) {
			err := os.MkdirAll(path, 0777)
			if err != nil {
				logger.Fatal(err)
			}
		}

		// create log file
		var filePath = path + "/" + viper.GetString("LOG_FILE")
		var file, err = os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			logger.Fatal(err)
		}
		logger.SetOutput(file)
	}

	return logger
}
