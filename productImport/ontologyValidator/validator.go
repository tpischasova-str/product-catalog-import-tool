package ontologyValidator

import (
	"ts/logger"
	"ts/productImport/attribute"
	"ts/productImport/mapping"
	"ts/productImport/ontologyRead/models"
	"ts/productImport/product"
	"ts/productImport/reports"
)

type Validator struct {
	logger           logger.LoggerInterface
	productHandler   product.ProductHandlerInterface
	ColumnMap        *ColumnMap
	uomMappingConfig *mapping.UoMMapConfig
}

type ColumnMap struct {
	Category  string
	ProductID string
	Name      string
	UOM       string
}

func NewValidator(deps Deps) ValidatorInterface {
	m := deps.Mapper.GetColumnMapConfig()
	return &Validator{
		logger:         deps.Logger,
		productHandler: deps.ProductHandler,
		ColumnMap: &ColumnMap{
			Category:  m.Category,
			ProductID: m.ProductID,
			Name:      m.Name,
			UOM:       m.UOM,
		},
		uomMappingConfig: deps.Mapper.GetUoMMapConfig(),
	}
}

func (v *Validator) InitialValidation(
	mapping map[string]string,
	rules *models.OntologyConfig,
	sourceData []map[string]interface{}) ([]reports.Report, bool) {
	report, isErr := v.validateProductsAgainstRules(mapping,
		rules,
		sourceData,
	)
	return report, isErr
}

func (v *Validator) SecondaryValidation(
	mapping map[string]string,
	rules *models.OntologyConfig,
	sourceData []map[string]interface{},
	attributeData []*attribute.Attribute,
) ([]reports.Report, bool) {

	parsedProducts := v.productHandler.InitParsedSourceData(sourceData)
	if attributeData != nil && len(attributeData) > 0 {
		report, isErr := v.validateAttributesAgainstRules(rules, parsedProducts, attributeData)
		return report, isErr
	}

	report, isErr := v.validateProductsAgainstRules(mapping,
		rules,
		sourceData,
	)
	return report, isErr
}
