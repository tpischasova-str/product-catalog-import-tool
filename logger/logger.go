package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
)

const (
	defaultLevel = "info"
)

type Logger struct {
	logLevel  logrus.Level
	formatter logrus.Formatter
}

func NewLogger(deps Deps) LoggerInterface {
	level, levelInitErr := initLevel(deps.Config.Service.LogLevel)
	logger := Logger{
		logLevel:  level,
		formatter: initFormatter(),
	}

	logrus.SetLevel(logger.logLevel)
	logrus.SetFormatter(logger.formatter)

	if levelInitErr != nil {
		logger.Warn(
			fmt.Sprintf("Failed to init logger with level from config, it was inited with default '%v' level", logger.logLevel),
			levelInitErr)
	}
	return &logger
}

func initFormatter() logrus.Formatter {
	customFormatter := new(logrus.TextFormatter)
	customFormatter.FullTimestamp = true
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	return customFormatter
}

func initLevel(configLevel string) (logrus.Level, error) {
	if configLevel == "" {
		parsedLevel, _ := logrus.ParseLevel(defaultLevel)
		return parsedLevel, nil
	} else {
		parsedLevel, err := logrus.ParseLevel(configLevel)
		if err != nil {
			parsedLevel, _ = logrus.ParseLevel(defaultLevel)
		}
		return parsedLevel, err
	}
}

func (l *Logger) Info(msg string) {
	logrus.WithTime(time.Now()).Info(msg)
}

func (l *Logger) Debug(msg string, fields map[string]interface{}) {
	logrus.WithTime(time.Now()).WithFields(fields).Debug(msg)
}

func (l *Logger) Warn(msg string, err error) {
	if err != nil {
		logrus.WithTime(time.Now()).WithError(err).Warn(msg)
	} else {
		logrus.WithTime(time.Now()).Warn(msg)
	}
}

func (l *Logger) Error(msg string, err error) {
	if err != nil {
		logrus.WithTime(time.Now()).WithError(err).Error(msg)
	} else {
		logrus.WithTime(time.Now()).Error(msg)
	}
}

func (l *Logger) Fatal(msg string, err error) {
	if err != nil {
		logrus.WithTime(time.Now()).WithError(err).Fatal(msg)
	} else {
		logrus.WithTime(time.Now()).Fatal(msg)
	}
}
