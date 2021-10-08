package ontologyValidator

import (
	"fmt"
	"ts/productImport/ontologyRead/models"
	"ts/utils"
)

func (v *Validator) isValidAttributeUoM(actualUom string, attributeRule *models.AttributeConfig) (bool, string) {
	if attributeRule.MeasurementUoM == "" {
		return true, ""
	}
	if utils.TrimAll(actualUom) != utils.TrimAll(attributeRule.MeasurementUoM) {
		return false, fmt.Sprintf("The attribute's UOM value should be '%v'", attributeRule.MeasurementUoM)
	}
	return true, ""
}
