package adapters

import (
	"go.uber.org/dig"
	"ts/config"
	"ts/logger"
)

type Deps struct{
	dig.In
	Config *config.Config
	Logger logger.LoggerInterface
}