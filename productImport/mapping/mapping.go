package mapping

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"ts/logger"
)

type mapping struct {
	logger    logger.LoggerInterface
	rawMap    map[string]string
	parsedMap *ColumnMapConfig
}

func NewMappingHandler(deps Deps) MappingHandlerInterface {
	rawMap := mapping{
		logger: deps.Logger,
	}
	rawMap.init(deps.Config.ProductCatalog.MappingPath)
	rawMap.parsedMap = rawMap.NewColumnMap(rawMap.rawMap)
	return &rawMap
}

func (m *mapping) init(path string) map[string]string {
	var rawColumnMap map[string]string
	if path != "" {
		if _, err := os.Stat(path); !os.IsNotExist(err) {
			m.upload(path)
			rawColumnMap = m.Get()
		}
	}
	return rawColumnMap
}

func (m *mapping) upload(mappingConfigPath string) {
	data, err := ioutil.ReadFile(mappingConfigPath)
	if err != nil {
		m.logger.Fatal(fmt.Sprintf("unable to load mapping file from %s", mappingConfigPath), err)
	}
	rawMapping := &RawMapping{}
	if err := yaml.Unmarshal(data, rawMapping); err != nil {
		m.logger.Fatal(fmt.Sprintf("unable to unmarshal mapping file %s", mappingConfigPath), err)
	}
	m.rawMap = rawMapping.ToConfig()
}

func (m *mapping) Get() map[string]string {
	return m.rawMap
}

func (m *mapping) GetColumnMapConfig() *ColumnMapConfig {
	return m.parsedMap
}
