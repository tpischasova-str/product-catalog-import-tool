package file

import (
	"fmt"
	"ts/file/csvFile"
	"ts/file/xlsxFile"
)

func XLSXToCSV(sourceFilePath string, sheet string, destinationFilePath string) error {
	sourcePath := xlsxFile.BuildPath(sourceFilePath, sheet)
	data, err := xlsxFile.Read(sourcePath)
	if err != nil {
		return fmt.Errorf("failed to convert XLSX to csv: %v", err)
	}
	err = csvFile.Write(destinationFilePath, data)
	if err != nil {
		return fmt.Errorf("failed to convert XLSX to csv: %v", err)
	}
	return nil
}
