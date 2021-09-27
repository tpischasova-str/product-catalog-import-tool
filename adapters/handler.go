package adapters

import (
	"fmt"
	"ts/adapters/csvH"
	"ts/adapters/excelH"
	"ts/adapters/txtH"
	"ts/logger"
)

type FileType string

const (
	CSV  FileType = "csv"
	XLSX FileType = "xlsx"
	TXT  FileType = "txt"
)

type Handler struct {
	logger    logger.LoggerInterface
	Adapter   AdapterInterface
	Delimiter rune
	header    []string
	LineChan  chan interface{}
}

func NewHandler(deps Deps) HandlerInterface {
	return &Handler{
		logger: deps.Logger,
	}
}

func (h *Handler) Init(t FileType) {
	switch t {
	case XLSX:
		h.Adapter = &excelH.Adapter{}
	case CSV:
		h.Adapter = &csvH.Adapter{}
	case TXT:
		h.Adapter = &txtH.Adapter{}
	default:
		h.logger.Fatal("unsupported source file type (only csv and xlsx are supported)", nil)
	}
}

func (h *Handler) GetHeader() []string {
	return h.header
}

func (h *Handler) Write(filepath string, data [][]string) {
	err := h.Adapter.Write(filepath, data)
	if err != nil {
		h.logger.Fatal(fmt.Sprintf("failed to write %v file", filepath), err)
	}
}

func (h *Handler) Parse(filePath string) []map[string]interface{} {
	res, err := h.Adapter.Parse(filePath)
	h.header = h.Adapter.GetHeader()
	if err != nil {
		h.logger.Fatal(fmt.Sprintf("failed to Read file %v", filePath), err)
	}
	return res
}
