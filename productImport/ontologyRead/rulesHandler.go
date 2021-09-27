package ontologyRead

import (
	"fmt"
	"os"
	"ts/adapters"
	"ts/logger"
	"ts/productImport/ontologyRead/models"
	"ts/productImport/ontologyRead/rawOntology"
)

type RulesHandler struct {
	logger     logger.LoggerInterface
	sourcePath string
	reader     adapters.HandlerInterface
}

func NewRulesHandler(deps Deps) *RulesHandler {
	return &RulesHandler{
		logger:     deps.Logger,
		sourcePath: deps.Config.ProductCatalog.OntologyPath,
		reader:     deps.Handler,
	}
}

func (h *RulesHandler) InitRulesConfig() (*models.OntologyConfig, error) {
	var rulesConfig *models.OntologyConfig
	var rules *rawOntology.RawOntology
	sourcePath := h.sourcePath
	if sourcePath != "" {
		if _, err := os.Stat(sourcePath); !os.IsNotExist(err) {
			rules = h.UploadRules(sourcePath)
			rulesConfig = rules.ToConfig()
		} else {
			h.logger.Fatal(fmt.Sprintf("ontology file does not exists. Please fill and add it to %v", sourcePath), nil)
		}
	} else {
		h.logger.Fatal("ontology path is not specified", nil)
	}
	return rulesConfig, nil
}

func (h *RulesHandler) UploadRules(path string) *rawOntology.RawOntology {

	ext := adapters.GetFileType(path)
	h.reader.Init(ext)
	parsedRaws := h.reader.Parse(path)
	actualHeader := h.reader.GetHeader()
	header, err := processHeader(actualHeader)
	if err != nil {
		h.logger.Fatal("failed to upload rules", err)
	}
	o := h.processOntology(parsedRaws, header)
	h.logger.Info(fmt.Sprintf("Rules upload finished. Proceeded %v lines, uploaded %v categories", len(parsedRaws)+1, o.GetCategoriesCount()))
	return o
}

func processHeader(parsedHeader []string) (*rawOntology.RawHeader, error) {
	resHeader := rawOntology.NewHeader(parsedHeader)
	if err := resHeader.ValidateHeader(); err != nil {
		return nil, err
	}
	return resHeader, nil
}

func (h *RulesHandler) processOntology(parsedRaws []map[string]interface{}, header *rawOntology.RawHeader) *rawOntology.RawOntology {
	o := rawOntology.NewRawOntology()

	for i, raw := range parsedRaws {
		var errors []error

		errors = rawOntology.ValidateRaw(raw, header)
		if len(errors) == 0 {
			rawAttribute := rawOntology.NewRawAttribute(raw, header)
			rawCategory := rawOntology.NewRawCategory(raw, header)
			err := o.AddCategoryAttribute(rawCategory, rawAttribute)
			if err != nil {
				h.logger.Debug(fmt.Sprintf("raw %v error", i),
					map[string]interface{}{"error": err})
			}
		} else {
			h.logger.Warn(fmt.Sprintf("raw %v: validation errors: %v", i, errors), nil)
		}
	}
	return o
}
