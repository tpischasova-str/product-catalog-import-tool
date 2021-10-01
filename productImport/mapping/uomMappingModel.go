package mapping

import (
	"ts/utils"
)

type UoMItem struct {
	DefaultKey string
	MappedKey  string
}

type UoMMapConfig struct {
	items map[string]*UoMItem
}

func NewUoMMapConfig(items []*UoMItem) *UoMMapConfig {
	res := make(map[string]*UoMItem)
	for _, item := range items {
		res[utils.TrimAll(item.MappedKey)] = item
	}
	return &UoMMapConfig{
		items: res,
	}
}

func (u *UoMMapConfig) GetDefaultUoMValueByMapped(value string) string {
	if res, ok := u.items[utils.TrimAll(value)]; ok {
		return res.DefaultKey
	}
	return ""
}

func (u *UoMMapConfig) GetActualUoMValueByDefault(defaultValue string) string {
	for _, u := range u.items {
		if utils.TrimAll(defaultValue) == utils.TrimAll(u.DefaultKey) {
			return u.MappedKey
		}
	}
	return ""
}
