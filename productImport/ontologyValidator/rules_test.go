package ontologyValidator

import (
	"testing"
	"ts/logger"
	"ts/productImport/mapping"
	"ts/productImport/ontologyRead/models"
	"ts/productImport/product"
)

func TestValidator_isValidAttributeUoM(t *testing.T) {
	type fields struct {
		logger           logger.LoggerInterface
		productHandler   product.ProductHandlerInterface
		ColumnMap        *ColumnMap
		uomMappingConfig *mapping.UoMMapConfig
	}
	type args struct {
		actualUom     string
		attributeRule *models.AttributeConfig
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
		want1  string
	}{
		{
			name: "Unmapped UOM considered as valid if it is equal to UOM from rules",
			fields: fields{
				ColumnMap: &ColumnMap{
					UOM: "UOM",
				},
				uomMappingConfig: &mapping.UoMMapConfig{},
			},
			args: args{
				actualUom: "kg",
				attributeRule: &models.AttributeConfig{
					MeasurementUoM: "Kg",
				},
			},
			want:  true,
			want1: "",
		},
		{
			name: "Mapped UOM considered as valid if it's default key is equal to UOM from rules",
			fields: fields{
				ColumnMap: &ColumnMap{
					UOM: "UOM",
				},
				uomMappingConfig: &mapping.UoMMapConfig{
					Items: map[string]*mapping.UoMItem{
						"kilo": {
							DefaultKey: "KG",
							MappedKey:  "KILO",
						},
					},
				},
			},
			args: args{
				actualUom: "Kilo",
				attributeRule: &models.AttributeConfig{
					MeasurementUoM: "Kg",
				},
			},
			want:  true,
			want1: "",
		},
		{
			name: "Empty UOM considered as valid if UoM in rules was not defined",
			fields: fields{
				ColumnMap: &ColumnMap{
					UOM: "UOM",
				},
				uomMappingConfig: &mapping.UoMMapConfig{},
			},
			args: args{
				actualUom: "",
				attributeRule: &models.AttributeConfig{
					MeasurementUoM: "",
				},
			},
			want:  true,
			want1: "",
		},
		{
			name: "Unmapped UOM considered as invalid if it is not equal to UOM from rules",
			fields: fields{
				ColumnMap: &ColumnMap{
					UOM: "UOM",
				},
				uomMappingConfig: &mapping.UoMMapConfig{},
			},
			args: args{
				actualUom: "pound",
				attributeRule: &models.AttributeConfig{
					MeasurementUoM: "Kg",
				},
			},
			want:  false,
			want1: "The attribute's UOM value should be 'Kg'",
		},
		{
			name: "Mapped UOM considered as invalid if it's default key is not equal to UOM from rules",
			fields: fields{
				ColumnMap: &ColumnMap{
					UOM: "UOM",
				},
				uomMappingConfig: &mapping.UoMMapConfig{
					Items: map[string]*mapping.UoMItem{
						"pounds": {
							DefaultKey: "PN",
							MappedKey:  "Pounds",
						},
					},
				},
			},
			args: args{
				actualUom: "Pounds",
				attributeRule: &models.AttributeConfig{
					MeasurementUoM: "Kg",
				},
			},
			want:  false,
			want1: "The attribute's UOM value should be 'Kg'",
		},
		{
			name: "Empty UOM considered as invalid if UoM in rules was defined",
			fields: fields{
				ColumnMap: &ColumnMap{
					UOM: "UOM",
				},
				uomMappingConfig: &mapping.UoMMapConfig{},
			},
			args: args{
				actualUom: "",
				attributeRule: &models.AttributeConfig{
					MeasurementUoM: "Kg",
				},
			},
			want:  false,
			want1: "The attribute's UOM value should be 'Kg'",
		},
		{
			name: "Not empty UOM considered as valid if UoM in rules was not defined",
			fields: fields{
				ColumnMap: &ColumnMap{
					UOM: "UOM",
				},
				uomMappingConfig: &mapping.UoMMapConfig{},
			},
			args: args{
				actualUom: "KG",
				attributeRule: &models.AttributeConfig{
					MeasurementUoM: "",
				},
			},
			want:  true,
			want1: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Validator{
				logger:           tt.fields.logger,
				productHandler:   tt.fields.productHandler,
				ColumnMap:        tt.fields.ColumnMap,
				uomMappingConfig: tt.fields.uomMappingConfig,
			}
			got, got1 := v.isValidAttributeUoM(tt.args.actualUom, tt.args.attributeRule)
			if got != tt.want {
				t.Errorf("isValidAttributeUoM() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("isValidAttributeUoM() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
