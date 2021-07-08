package sheetConverter

import (
	"fmt"
	"ts/adapters"
	"ts/file"
)

type FailuresConverter struct { //todo name
	sheet           string
	destinationPath string
}

func NewFailuresConverter(sheet string, destinationPath string) *FailuresConverter {
	return &FailuresConverter{
		sheet:           sheet,
		destinationPath: destinationPath,
	}
}

func (c *FailuresConverter) Convert(filePath string) error {
	destinationPath := buildFailuresPath(filePath, c.destinationPath)
	err := file.XLSXToCSV(filePath, c.sheet, destinationPath)
	return err
}

func buildFailuresPath(filePath string, catalogPath string) string {
	fileName := adapters.GetFileName(filePath)
	return fmt.Sprintf(
		"%v%v-failures.%v",
		adapters.AddSlashToPath(catalogPath),
		fileName,
		adapters.CSV)
}
