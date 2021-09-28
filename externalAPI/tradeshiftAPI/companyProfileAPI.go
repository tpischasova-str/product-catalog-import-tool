package tradeshiftAPI

import (
	"fmt"
)

func (t *TradeshiftAPI) GetBuyer(buyerID string) (map[string]interface{}, error) {
	method := fmt.Sprintf("/company-profile/v0/company-card/companies/%v", buyerID)

	resp, err := t.Client.Get(method, nil)

	if err != nil {
		return nil, err
	}
	r, err := t.Client.ParseResponse(resp)
	return r, err
}
