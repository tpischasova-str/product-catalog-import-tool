package mapping

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"ts/logger"
)

type mapping struct {
	logger            logger.LoggerInterface
	uomMappingPath    string
	columnMappingPath string
	rawMap            map[string]string
	parsedMap         *ColumnMapConfig
	uomMap            *UoMMapConfig
}

func NewMappingHandler(deps Deps) MappingHandlerInterface {
	conf := deps.Config.ProductCatalog
	rawMap := mapping{
		logger:            deps.Logger,
		columnMappingPath: conf.MappingPath,
		uomMappingPath:    conf.UoMMappingPath,
	}
	rawMap.initProductsMapping()
	rawMap.initUoMapping()
	rawMap.parsedMap = rawMap.NewColumnMap(rawMap.rawMap)
	return &rawMap
}

func (m *mapping) initProductsMapping() map[string]string {
	var rawColumnMap map[string]string
	if m.columnMappingPath != "" {
		if _, err := os.Stat(m.columnMappingPath); !os.IsNotExist(err) {
			m.readColumnMapping(m.columnMappingPath)
			rawColumnMap = m.Get()
		}
	}
	return rawColumnMap
}

func (m *mapping) readColumnMapping(mappingConfigPath string) {
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

func (m *mapping) initUoMapping() {
	res := make([]*UoMItem, 0)
	if m.uomMappingPath != "" {
		if _, err := os.Stat(m.uomMappingPath); !os.IsNotExist(err) {
			res = m.readUoMMapping(m.uomMappingPath)
		}
	}
	m.uomMap = NewUoMMapConfig(res)
}

func (m *mapping) readUoMMapping(uomPath string) []*UoMItem {
	data, err := ioutil.ReadFile(uomPath)
	if err != nil {
		m.logger.Fatal(fmt.Sprintf("unable to load mapping file from %s", uomPath), err)
	}
	rawUoMMapping := &RawUomMapping{}
	if err := yaml.Unmarshal(data, rawUoMMapping); err != nil {
		m.logger.Fatal(fmt.Sprintf("unable to unmarshal mapping file %s", uomPath), err)
	}
	return rawUoMMapping.ToConfig()
}

func (m *mapping) Get() map[string]string {
	return m.rawMap
}

func (m *mapping) GetColumnMapConfig() *ColumnMapConfig {
	return m.parsedMap
}

func (m *mapping) GetUoMMapConfig() *UoMMapConfig {
	return m.uomMap
}
