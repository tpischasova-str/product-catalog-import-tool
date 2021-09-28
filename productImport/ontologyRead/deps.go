package ontologyRead

import (
	"go.uber.org/dig"
	"ts/adapters"
	"ts/config"
	"ts/logger"
)

type Deps struct {
	dig.In
	Config       *config.Config
	Logger       logger.LoggerInterface
	Handler      adapters.HandlerInterface
	FilesManager *adapters.FileManager
}
