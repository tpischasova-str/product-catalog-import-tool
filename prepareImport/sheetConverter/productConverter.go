package sheetConverter

import (
	"fmt"
	"ts/adapters"
	"ts/file"
)

type ProductConverter struct {
	sheet           string
	destinationPath string
}

func NewProductConverter(sheet string, destinationPath string) *ProductConverter {
	return &ProductConverter{
		sheet:           sheet,
		destinationPath: destinationPath,
	}
}

func (c *ProductConverter) Convert(filePath string) error {
	destinationPath := buildProductPath(filePath, c.destinationPath)
	err := file.XLSXToCSV(filePath, c.sheet, destinationPath)
	return err
}

func buildProductPath(filePath string, catalogPath string) string {
	return fmt.Sprintf(
		"%v%v.%v",
		adapters.AddSlashToPath(catalogPath),
		adapters.GetFileName(filePath),
		adapters.CSV)
}
