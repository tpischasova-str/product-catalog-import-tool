package logger

import (
	"go.uber.org/dig"
	"ts/config"
)

type Deps struct {
	dig.In
	Config *config.Config
}

type LoggerInterface interface {
	Debug(msg string, fields map[string]interface{})
	Info(msg string)
	Warn(msg string, err error)
	Error(msg string, err error)
	Fatal(msg string, err error)
}
