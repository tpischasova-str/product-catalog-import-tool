package ontologyValidator

import (
	"testing"
	"ts/logger"
	"ts/productImport/mapping"
	"ts/productImport/product"
)

func TestValidator_buildUomColumn(t *testing.T) {
	type fields struct {
		logger           logger.LoggerInterface
		productHandler   product.ProductHandlerInterface
		ColumnMap        *ColumnMap
		uomMappingConfig *mapping.UoMMapConfig
	}
	type args struct {
		mapping  map[string]string
		attrName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "mapped value should be returned if column has mapping",
			fields: fields{},
			args: args{
				mapping: map[string]string{
					"Default": "Actual",
				},
				attrName: "Actual",
			},
			want: "Actual_UOM",
		},

		{
			name:   "unmapped value should be returned if there is no mapping for column",
			fields: fields{},
			args: args{
				mapping: map[string]string{
					"Seno": "Soloma",
				},
				attrName: "Actual",
			},
			want: "Actual_UOM",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Validator{
				logger:           tt.fields.logger,
				productHandler:   tt.fields.productHandler,
				ColumnMap:        tt.fields.ColumnMap,
				uomMappingConfig: tt.fields.uomMappingConfig,
			}
			if got := v.buildUomColumn(tt.args.mapping, tt.args.attrName); got != tt.want {
				t.Errorf("buildUomColumn() = %v, want %v", got, tt.want)
			}
		})
	}
}
