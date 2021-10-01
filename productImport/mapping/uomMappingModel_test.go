package mapping

import "testing"

func TestUoMMapConfig_GetActualUoMValueByDefault(t *testing.T) {
	type fields struct {
		items map[string]*UoMItem
	}
	type args struct {
		defaultValue string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "Should be able to find actual UoM name in mapping by UBL key",
			fields: fields{
				items: map[string]*UoMItem{
					"actual1": {
						DefaultKey: "Act1",
						MappedKey:  "Actual1",
					},
				},
			},
			args: args{
				"Act1",
			},
			want: "Actual1",
		},
		{
			name: "Should be empty result if there is no mapping for UBL key",
			fields: fields{
				items: map[string]*UoMItem{
					"actual1": {
						DefaultKey: "Act1",
						MappedKey:  "Actual1",
					},
				},
			},
			args: args{
				"Key1",
			},
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UoMMapConfig{
				items: tt.fields.items,
			}
			if got := u.GetActualUoMValueByDefault(tt.args.defaultValue); got != tt.want {
				t.Errorf("GetActualUoMValueByDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUoMMapConfig_GetDefaultUoMValueByMapped(t *testing.T) {
	type fields struct {
		items map[string]*UoMItem
	}
	type args struct {
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "Should be found UBL-formatted key by actual UoM name",
			fields: fields{
				items: map[string]*UoMItem{
					"act1": {
						DefaultKey: "Actual 1",
						MappedKey:  "Act 1",
					},
				},
			},
			args: args{
				value: "Act1",
			},
			want: "Actual 1",
		},
		{
			name: "Should be empty result if there is no known mapped value for UoM name",
			fields: fields{
				items: map[string]*UoMItem{
					"act1": {
						DefaultKey: "Actual 1",
						MappedKey:  "Act 1",
					},
				},
			},
			args: args{
				value: "Key1",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UoMMapConfig{
				items: tt.fields.items,
			}
			if got := u.GetDefaultUoMValueByMapped(tt.args.value); got != tt.want {
				t.Errorf("GetDefaultUoMValueByMapped() = %v, want %v", got, tt.want)
			}
		})
	}
}
