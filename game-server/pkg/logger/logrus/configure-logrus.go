package configureLogrus

import (
	"os"

	"github.com/sirupsen/logrus"
)

const (
	JSON = "JSON"
	TEXT = "TEXT"
)

type LogrusConfig struct {
	// text || json
	OutputType      string
	LogLevel        string
	SetReportCaller bool
}

func SetUpLogrus(config *LogrusConfig) *logrus.Logger {

	logger := logrus.New()

	outputFormat := getOutputType(config.OutputType)
	logLevel := getLogLevel(config.LogLevel)

	logger.SetFormatter(outputFormat)
	logger.SetLevel(logLevel)
	//Добавить вариации оутпутов
	logger.SetOutput(os.Stdout)
	logger.SetReportCaller(config.SetReportCaller)

	return logger

}

func getOutputType(outputType string) logrus.Formatter {

	switch outputType {
	case JSON:
		return new(logrus.JSONFormatter)

	default:
		return new(logrus.TextFormatter)

	}

}

func getLogLevel(logLevel string) logrus.Level {

	switch logLevel {
	case "DEBUG":
		return logrus.DebugLevel

	case "ERROR":
		return logrus.ErrorLevel

	case "WARN":
		return logrus.WarnLevel

	default:
		return logrus.InfoLevel
	}

}
