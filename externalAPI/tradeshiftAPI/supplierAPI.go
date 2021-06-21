package tradeshiftAPI

import (
	"fmt"
	"net/url"
	"strconv"
	"ts/externalAPI/rest"
)

func (t *TradeshiftAPI) GetIdentifier() (map[string]interface{}, error) {
	method := "/product-engine/supplier/supplier/v1/properties/identifier"
	resp, err := t.Client.Get(method, nil)
	if err != nil {
		return nil, err
	}
	r, err := rest.ParseResponse(resp)
	return r, err
}

func (t *TradeshiftAPI) SetIdentifier(identifier string) error {
	method := "/product-engine/supplier/supplier/v1/properties/identifier"
	data := map[string]interface{}{
		"autoGenerated": false,
		"name":          identifier,
	}
	_, err := t.Client.Post(method, rest.BuildBody(data), nil)
	return err
}

func (t *TradeshiftAPI) UploadFile(filePath string) (map[string]interface{}, error) {
	method := "/product-engine/supplier/supplier/v1/files"

	resp, err := t.Client.PostFile(method, filePath)
	r, err := rest.ParseResponse(resp)
	return r, err
}

func (t *TradeshiftAPI) RunImportAction(fileID string) (string, error) {
	method := fmt.Sprintf("/product-engine/supplier/supplier/v1/product-import/files/%v/actions/import-products", url.QueryEscape(fileID))
	resp, err := t.Client.Post(
		method,
		nil,
		[]rest.UrlParam{
			{
				Key:   "currency",
				Value: "USD",
			},
			{
				Key:   "fileLocale",
				Value: "en_US",
			},
		})
	if err != nil {
		return "", err
	}
	r, err := rest.ParseResponseToString(resp)
	return r, err
}

func (t *TradeshiftAPI) GetActionResult(actionID string) (map[string]interface{}, error) {
	method := fmt.Sprintf("/product-engine/supplier/supplier/v1/actions/%v", url.QueryEscape(actionID))
	resp, err := t.Client.Get(method, nil)
	r, err := rest.ParseResponse(resp)
	return r, err
}

func (t *TradeshiftAPI) GetImportResult(actionID string) (string, error) {
	method := fmt.Sprintf("/product-engine/supplier/supplier/v1/actions/%v/reports/import-product-report/download", url.QueryEscape(actionID))
	resp, err := t.Client.Get(method, nil)
	r, err := rest.ParseResponseToString(resp)
	return r, err
}

func (t *TradeshiftAPI) SearchOffer(name string) (map[string]interface{}, error) {
	method := "/product-engine/supplier/supplier/v1/offers"
	limit := 5

	params := []rest.UrlParam{
		{
			Key:   "advancedSearch",
			Value: buildAdvancedSearchValue(name),
		},
		{
			Key:   "sort",
			Value: "name",
		},
		{
			Key:   "limit",
			Value: strconv.Itoa(limit),
		},
	}
	resp, err := t.Client.Get(method, params)
	r, err := rest.ParseResponse(resp)
	return r, err
}

func buildAdvancedSearchValue(name string) string {
	return fmt.Sprintf("{\"name\":\"%v\"}", url.QueryEscape(name))
}

func (t *TradeshiftAPI) CreateOffer(name string, buyerID string) (string, error) {
	method := "/product-engine/supplier/supplier/v1/offers"
	params := []rest.UrlParam{
		{
			Key:   "buyerId",
			Value: url.QueryEscape(buyerID),
		},
	}

	resp, err := t.Client.Post(
		method,
		rest.BuildBody(name),
		params)

	if err != nil {
		return "", err
	}
	r, err := rest.ParseResponseToString(resp)
	if err != nil {
		return "", err
	}
	return r, err
}
