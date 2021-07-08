package sheetConverter

import (
	"fmt"
	"ts/adapters"
	"ts/file"
)

type OfferConverter struct {
	sheet           string
	destinationPath string
}

func NewOfferConverter(sheet string, destinationPath string) *OfferConverter {
	return &OfferConverter{
		sheet:           sheet,
		destinationPath: destinationPath,
	}
}

func (c *OfferConverter) Convert(filePath string) error {
	destinationPath := buildOfferPath(filePath, c.destinationPath)
	err := file.XLSXToCSV(filePath, c.sheet, destinationPath)
	return err
}

func buildOfferPath(filePath string, catalogPath string) string {
	fileName := adapters.GetFileName(filePath)
	return fmt.Sprintf(
		"%v%v.%v",
		adapters.AddSlashToPath(catalogPath),
		fileName,
		adapters.CSV)
}
