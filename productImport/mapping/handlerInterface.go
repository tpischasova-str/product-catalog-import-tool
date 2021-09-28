package mapping

import (
	"go.uber.org/dig"
	"ts/config"
	"ts/logger"
)

type MappingHandlerInterface interface {
	init(mappingConfigPath string) map[string]string
	Get() map[string]string
	GetColumnMapConfig() *ColumnMapConfig
}

type Deps struct {
	dig.In
	Config *config.Config
	Logger logger.LoggerInterface
}
