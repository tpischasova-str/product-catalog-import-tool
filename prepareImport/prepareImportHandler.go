package prepareImport

import (
	"fmt"
	"go.uber.org/dig"
	"path/filepath"
	"strings"
	"ts/adapters"
	"ts/config"
	"ts/logger"
)

type Handler struct {
	sourcePath          string
	sentPath            string
	shouldBeSkipped     bool
	productConverter    *XLSXSheetToCSVConverter
	attributesConverter *XLSXSheetToCSVConverter
	offerConverter      *XLSXSheetToCSVConverter
	offerItemConverter  *XLSXSheetToCSVConverter
	logger              logger.LoggerInterface
}

type Deps struct {
	dig.In
	Config *config.Config
	Logger logger.LoggerInterface
}

func NewPrepareImportHandler(deps Deps) *Handler {
	conf := deps.Config
	xlsxConfig := deps.Config.XLSXConfig
	if deps.Config.XLSXConfig == nil {
		return nil
	}
	sheetsConf := xlsxConfig.Sheets
	return &Handler{
		logger:     deps.Logger,
		sourcePath: xlsxConfig.SourcePath,
		sentPath:   xlsxConfig.SentPath,
		productConverter: NewXLSXSheetToCSVConverter(
			sheetsConf.Products.Name,
			sheetsConf.Products.HeaderRowsToSkip,
			conf.ProductCatalog.InProgressPath,
			""),
		attributesConverter: NewXLSXSheetToCSVConverter(
			sheetsConf.Attributes.Name,
			sheetsConf.Attributes.HeaderRowsToSkip,
			conf.ProductCatalog.SourcePath,
			"_attributes"),
		offerConverter: NewXLSXSheetToCSVConverter(
			sheetsConf.Offers.Name,
			sheetsConf.Offers.HeaderRowsToSkip,
			conf.OfferCatalog.SourcePath,
			"_offers"),
		offerItemConverter: NewXLSXSheetToCSVConverter(
			sheetsConf.OfferItems.Name,
			sheetsConf.OfferItems.HeaderRowsToSkip,
			conf.OfferItemCatalog.SourcePath,
			"_offer_items"),
	}
}

func (h *Handler) Run() {
	files := getXLSXFiles(h.sourcePath)
	if len(files) == 0 {
		return
	}
	for _, fileName := range files {
		filePath := filepath.Join(
			h.sourcePath,
			fileName)
		err := h.convertSheetsData(filePath)
		if err != nil {
			h.logger.Error(fmt.Sprintf("failed to convert file %v", filePath), err)
		}
		_, err = adapters.MoveToPath(filePath, h.sentPath)
		if err != nil {
			h.logger.Error(fmt.Sprintf("failed to move %v to %v", filePath, h.sentPath), err)
		}
	}
}

func getXLSXFiles(path string) []string {
	var res []string
	files := adapters.GetFiles(path)
	for _, filePath := range files {
		if isXLSX(filePath) {
			res = append(res, filePath)
		}
	}
	return res
}

func isXLSX(filePath string) bool {
	res := strings.HasSuffix(strings.ToLower(filePath), ".xls") || strings.HasSuffix(strings.ToLower(filePath), ".xlsx")
	return res
}

func (h *Handler) convertSheetsData(filePath string) error {
	err := h.productConverter.Convert(filePath)
	if err != nil {
		return fmt.Errorf("failed to convert Products: %v", err)
	}
	err = h.offerConverter.Convert(filePath)
	if err != nil {
		return fmt.Errorf("failed to convert Offers: %v", err)
	}
	err = h.attributesConverter.Convert(filePath)
	if err != nil {
		return fmt.Errorf("failed to convert Attributes: %v", err)
	}

	err = h.offerItemConverter.Convert(filePath)
	if err != nil {
		return fmt.Errorf("failed to convert Offer Items: %v", err)
	}
	return nil
}
