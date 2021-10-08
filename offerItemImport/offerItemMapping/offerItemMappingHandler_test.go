package offerItemMapping

import (
	"reflect"
	"testing"
	"ts/productImport/mapping"
)

func TestOfferItemMappingHandler_buildHeader(t *testing.T) {
	type fields struct {
		columnMap         *mapping.ColumnMapConfig
		sourcePath        string
		successReportPath string
	}
	type args struct {
		headerRow []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		{
			name: "positive: mapped columnn names should be converted to default values.",
			fields: fields{
				columnMap: &mapping.ColumnMapConfig{
					ProductID: "SKU",
					Category:  "UNSPSC",
					OtherColumns: []*mapping.ColumnItem{
						{
							DefaultKey: "DefaultKey1",
							MappedKey:  "MappedKey1",
						},
						{
							DefaultKey: "DefaultKey2",
							MappedKey:  "MappedKey2",
						},
					},
				},
			},
			args: args{
				headerRow: []string{
					"Sku",
					"UNSPSC",
					"DefaultKey2",
				},
			},
			want: []string{
				"ID",
				"Category",
				"DefaultKey2",
			},
		},
		{
			name: "positive: if column name is not mapped, should be taken original column name",
			fields: fields{
				columnMap: &mapping.ColumnMapConfig{
					ProductID: "SKU",
					Category:  "UNSPSC",
					OtherColumns: []*mapping.ColumnItem{
						{
							DefaultKey: "DefaultKey1",
							MappedKey:  "MappedKey1",
						},
					},
				},
			},
			args: args{
				headerRow: []string{
					"Sku",
					"UNSPSC",
					"DefaultKey1",
					"Not Mapped Key",
				},
			},
			want: []string{
				"ID",
				"Category",
				"DefaultKey1",
				"Not Mapped Key",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oi := &OfferItemMappingHandler{
				columnMap:         tt.fields.columnMap,
				sourcePath:        tt.fields.sourcePath,
				successReportPath: tt.fields.successReportPath,
			}
			if got := oi.buildHeader(tt.args.headerRow); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildMappedHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}
