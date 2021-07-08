package prepareImport

import (
	"fmt"
	"go.uber.org/dig"
	"log"
	"ts/adapters"
	"ts/config"
	"ts/prepareImport/sheetConverter"
)

type Handler struct {
	sourcePath        string
	sentPath          string
	productConverter  *sheetConverter.ProductConverter
	failuresConverter *sheetConverter.FailuresConverter
	offerConverter    *sheetConverter.OfferConverter
}

type Deps struct {
	dig.In
	Config *config.Config
}

func NewPrepareImportHandler(deps Deps) *Handler {
	conf := deps.Config
	commonConf := deps.Config.CommonConfig
	return &Handler{
		sourcePath: commonConf.SourcePath,
		sentPath:   commonConf.SentPath,
		productConverter: sheetConverter.NewProductConverter(
			commonConf.Sheet.Products,
			conf.ProductCatalog.InProgressPath),
		failuresConverter: sheetConverter.NewFailuresConverter(
			commonConf.Sheet.Failures,
			conf.ProductCatalog.SourcePath),
		offerConverter: sheetConverter.NewOfferConverter(
			commonConf.Sheet.Offers,
			conf.OfferCatalog.SourcePath),
	}
}

func (h *Handler) Run() {
	files := adapters.GetXLSXFiles(h.sourcePath)
	if len(files) == 0 {
		log.Printf("no xlsx files for imports preparation are specified in %v", h.sourcePath)
		return
	}

	for _, fileName := range files {
		filePath := h.buildSourceFilePath(fileName)
		err := h.processFile(filePath)
		if err != nil {
			log.Printf("failed to process file %v: %v", filePath, err)
		}
		_, err = adapters.MoveToPath(filePath, h.sentPath)
		if err != nil {
			log.Printf("failed to move %v to %v: %v", filePath, h.sentPath, err)
		}
	}
}

func (h *Handler) buildSourceFilePath(fileName string) string {
	return fmt.Sprintf("%v%v", h.sourcePath, fileName)
}

func (h *Handler) processFile(filePath string) error {
	err := h.productConverter.Convert(filePath)
	if err != nil {
		return fmt.Errorf("failed to convert Products: %v", err)
	}
	err = h.offerConverter.Convert(filePath)
	if err != nil {
		return fmt.Errorf("failed to convert Offers: %v", err)
	}
	err = h.failuresConverter.Convert(filePath)
	if err != nil {
		return fmt.Errorf("failed to convert Attributes: %v", err)
	}
	return nil
}
