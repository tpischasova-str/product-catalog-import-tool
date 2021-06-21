package tradeshiftAPI

import (
	"testing"
)

func Test_buildAdvancedSearchValue(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "positive: with quote",
			args: args{
				name: "\"test\"",
			},
			want: "{\"name\":\"%22test%22\"}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildAdvancedSearchValue(tt.args.name); got != tt.want {
				t.Errorf("buildAdvancedSearchValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
