package mapping

type RawMapping struct {
	Map map[string]string `yaml:"column-mappings" validate:"required"`
}

func (rc *RawMapping) ToConfig() map[string]string {
	return rc.Map
}

type RawUomMapping struct {
	Map map[string]string `yaml:"uom-mappings"`
}

func (ru *RawUomMapping) ToConfig() []*UoMItem {
	res := make([]*UoMItem, 0)
	for key, value := range ru.Map {
		res = append(res, &UoMItem{
			DefaultKey: key,
			MappedKey:  value,
		})

	}
	return res
}
