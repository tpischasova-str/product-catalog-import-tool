package offerReader

import (
	"fmt"
	"go.uber.org/dig"
	"strings"
	"time"
	"ts/adapters"
	"ts/logger"
	"ts/utils"
)

const DateLayout = "2006-01-02"

type OfferReader struct {
	logger logger.LoggerInterface
	reader adapters.HandlerInterface
}

type RawOffer struct {
	Offer        string
	ReceiverName string
	Contract     string
	ValidFrom    time.Time
	ExpiresAt    time.Time
	Countries    []string
}

type Deps struct {
	dig.In
	Logger      logger.LoggerInterface
	Reader      adapters.HandlerInterface
	FileManager *adapters.FileManager
}

func NewOfferReader(deps Deps) *OfferReader {
	return &OfferReader{
		logger: deps.Logger,
		reader: deps.Reader,
	}
}

func (o *OfferReader) UploadOffers(path string) []RawOffer {
	ext := adapters.GetFileType(path)
	o.reader.Init(ext)
	parsedRaws := o.reader.Parse(path)
	actualHeader := o.reader.GetHeader()
	header, err := processHeader(actualHeader)
	if err != nil {
		o.logger.Error("failed to upload offers", err)
		return nil
	}
	or := o.processOffers(parsedRaws, header)
	return or
}

func (o *OfferReader) processOffers(raws []map[string]interface{}, header *RawHeader) []RawOffer {
	res := make([]RawOffer, len(raws))
	for i, item := range raws {
		offer, err := o.processOffer(header, item)
		if err != nil {
			o.logger.Error(fmt.Sprintf("An error occured while reading offer on line %v", i), err)
		} else {
			if offer != nil {
				res[i] = *offer
			}
		}
	}
	return res
}

func (o *OfferReader) processOffer(header *RawHeader, row map[string]interface{}) (*RawOffer, error) {
	if utils.IsEmptyMap(row) {
		return nil, nil
	}
	if row[header.Offer] == nil || row[header.Receiver] == nil {
		return nil, fmt.Errorf("row does not contain values in required columns (Offer, Receiver). Actual value: %v", row)
	}

	offer := RawOffer{
		Offer:        fmt.Sprintf("%v", row[header.Offer]),
		ReceiverName: fmt.Sprintf("%v", row[header.Receiver]),
	}
	if header.ContractID != "" && row[header.ContractID] != "" {
		offer.Contract = fmt.Sprintf("%v", row[header.ContractID])
	}
	if header.ValidFrom != "" && row[header.ValidFrom] != "" {
		date, err := time.Parse(DateLayout, fmt.Sprintf("%v", row[header.ValidFrom]))
		if err == nil {
			offer.ValidFrom = date
		} else {
			o.logger.Warn(fmt.Sprintf("invalid format of \"valid_from\" field: should be YYYY-MM-DD"), err)
		}
	}
	if header.ExpiresAt != "" && row[header.ExpiresAt] != "" {
		date, err := time.Parse(DateLayout, fmt.Sprintf("%v", row[header.ExpiresAt]))
		if err == nil {
			offer.ExpiresAt = date
		} else {
			o.logger.Warn("invalid format of \"expires_at\" field: should be YYYY-MM-DD: %v", err)
		}
	}
	if header.Countries != "" && row[header.Countries] != "" {
		offer.Countries = getCountries(row[header.Countries])
	}
	return &offer, nil
}

func getCountries(input interface{}) []string {
	return strings.SplitN(strings.ToLower(fmt.Sprintf("%v", input)), ",", -1)
}

func processHeader(parsedHeader []string) (*RawHeader, error) {
	resHeader := NewHeader(parsedHeader)
	if err := resHeader.ValidateHeader(); err != nil {
		return nil, err
	}
	return resHeader, nil
}
