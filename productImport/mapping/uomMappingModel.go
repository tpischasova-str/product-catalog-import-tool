package mapping

import (
	"ts/utils"
)

type UoMItem struct {
	DefaultKey string
	MappedKey  string
}

type UoMMapConfig struct {
	Items map[string]*UoMItem
}

func NewUoMMapConfig(items []*UoMItem) *UoMMapConfig {
	res := make(map[string]*UoMItem)
	for _, item := range items {
		res[utils.TrimAll(item.MappedKey)] = item
	}
	return &UoMMapConfig{
		Items: res,
	}
}

func (u *UoMMapConfig) GetDefaultUoMValueByMapped(value string) string {
	if res, ok := u.Items[utils.TrimAll(value)]; ok {
		return res.DefaultKey
	}
	return ""
}

func (u *UoMMapConfig) GetActualUoMValueByDefault(defaultValue string) string {
	for _, u := range u.Items {
		if utils.TrimAll(defaultValue) == utils.TrimAll(u.DefaultKey) {
			return u.MappedKey
		}
	}
	return ""
}
