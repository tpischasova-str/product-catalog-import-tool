package offerImport

import (
	"fmt"
	"go.uber.org/dig"
	"path/filepath"
	"ts/adapters"
	"ts/config"
	"ts/logger"
	"ts/offerImport/importHandler"
	"ts/offerImport/offerReader"
)

type OfferImportHandler struct {
	logger        logger.LoggerInterface
	sourcePath    string
	sentPath      string
	offerReader   *offerReader.OfferReader
	importHandler importHandler.ImportOfferInterface
}

type Deps struct {
	dig.In
	Config        *config.Config
	Logger        logger.LoggerInterface
	OfferReader   *offerReader.OfferReader
	ImportHandler importHandler.ImportOfferInterface
}

func NewOfferImportHandler(deps Deps) *OfferImportHandler {
	return &OfferImportHandler{
		logger:        deps.Logger,
		sourcePath:    deps.Config.OfferCatalog.SourcePath,
		sentPath:      deps.Config.OfferCatalog.SentPath,
		offerReader:   deps.OfferReader,
		importHandler: deps.ImportHandler,
	}
}

func (o *OfferImportHandler) RunCSV() {
	o.logger.Info("_________________________________")
	sourceFileNames := adapters.GetFiles(o.sourcePath)
	if len(sourceFileNames) == 0 {
		o.logger.Error("Source to import offers was not found. Skip step.", nil)
		return
	}

	for _, fileName := range sourceFileNames {
		o.runOfferImportFlow(fileName)
	}
}

func (o *OfferImportHandler) runOfferImportFlow(fileName string) {
	offers, err := o.uploadOffers(fileName)
	if err != nil {
		_, _ = adapters.MoveToPath(filepath.Join(o.sourcePath, fileName), o.sentPath)
		o.logger.Error(fmt.Sprintf("An error occurred while uploading the offers. Skip step, invalid file was moved to %v", o.sentPath), err)
		return
	}

	o.importHandler.ImportOffers(offers)
}

func (o *OfferImportHandler) uploadOffers(fileName string) ([]offerReader.RawOffer, error) {
	o.logger.Info(fmt.Sprintf("Offers file processing '%v' has been started", fileName))

	offers := o.offerReader.UploadOffers(filepath.Join(o.sourcePath, fileName))
	if len(offers) == 0 {
		return nil, fmt.Errorf(
			"0 offers were loaded from %v. Please, check file and try again",
			o.sourcePath)
	}
	err := o.processSourceFile(fileName)
	return offers, err
}

func (o *OfferImportHandler) processSourceFile(fileName string) error {
	_, err := adapters.MoveToPath(filepath.Join(o.sourcePath, fileName), o.sentPath)
	return err
}
