package importHandler

import (
	"go.uber.org/dig"
	"ts/config"
	"ts/externalAPI/tradeshiftAPI"
	"ts/logger"
	"ts/offerImport/offerReader"
)

type Status int

const (
	BuyerNotFound Status = 1
	OfferFound    Status = 2
	OfferCreated  Status = 4
	Failed        Status = 0
)

type Deps struct {
	dig.In
	Transport *tradeshiftAPI.TradeshiftAPI
	Config    *config.Config
	Logger    logger.LoggerInterface
}

type ImportOfferInterface interface {
	ImportOffers(offers []offerReader.RawOffer)
}
