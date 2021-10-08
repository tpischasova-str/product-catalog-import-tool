package tradeshiftImportHandler

import (
	"fmt"
	"go.uber.org/dig"
	"path/filepath"
	"ts/adapters"
	"ts/config"
	"ts/externalAPI/tradeshiftAPI"
	"ts/logger"
	"ts/outwardImport"
	"ts/outwardImport/importToTradeshift"
)

type TradeshiftHandler struct {
	logger               logger.LoggerInterface
	transport            *tradeshiftAPI.TradeshiftAPI
	filemanager          *adapters.FileManager
	handler              adapters.HandlerInterface
	outwardImportHandler outwardImport.OutwardImportInterface
	reportPath           string
}

type DepsH struct {
	dig.In
	Config               *config.Config
	Logger               logger.LoggerInterface
	TradeshiftAPI        *tradeshiftAPI.TradeshiftAPI
	FileManager          *adapters.FileManager
	FilesHandler         adapters.HandlerInterface
	OutwardImportHandler outwardImport.OutwardImportInterface
}

func NewTradeshiftHandler(deps DepsH) *TradeshiftHandler {
	h := deps.FilesHandler
	h.Init(adapters.TXT)

	return &TradeshiftHandler{
		logger:               deps.Logger,
		transport:            deps.TradeshiftAPI,
		filemanager:          deps.FileManager,
		handler:              h,
		outwardImportHandler: deps.OutwardImportHandler,
		reportPath:           deps.Config.ProductCatalog.ReportPath,
	}
}

func (th *TradeshiftHandler) ImportFeedToTradeshift(
	validationReportPath string) error {

	actionID, err := th.outwardImportHandler.ImportProducts(validationReportPath)
	if err != nil {
		return err
	}
	state, err := th.outwardImportHandler.WaitForImportComplete(actionID)
	if err != nil {
		return err
	}

	fileName := adapters.GetFileName(validationReportPath)

	err = th.outwardImportHandler.BuildProductAndOffersImportReport(
		actionID,
		filepath.Join(
			th.reportPath,
			fmt.Sprintf("report_products_%v", fileName)))
	if err != nil {
		return err
	}
	switch state {
	case importToTradeshift.CompleteImportState:
		th.logger.Info("Product import has been finished successfully")
	case importToTradeshift.CompleteWithErrorImportState:
		th.logger.Warn(fmt.Sprintf("Product import has been finished with errors. See report here '%v'", th.filemanager.ReportPath), nil)
	default:
		th.logger.Warn(fmt.Sprintf("Product import has been failed. See report here '%v'", th.filemanager.ReportPath), nil)
	}
	return nil
}
