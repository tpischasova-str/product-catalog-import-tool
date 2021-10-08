package mapping

import (
	"go.uber.org/dig"
	"ts/config"
	"ts/logger"
)

type MappingHandlerInterface interface {
	Get() map[string]string
	GetColumnMapConfig() *ColumnMapConfig
	GetUoMMapConfig() *UoMMapConfig
}

type Deps struct {
	dig.In
	Config *config.Config
	Logger logger.LoggerInterface
}
