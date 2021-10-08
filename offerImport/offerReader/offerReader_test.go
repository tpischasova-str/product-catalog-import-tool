package offerReader

import (
	"reflect"
	"testing"
	"ts/adapters"
	"ts/logger"
)


func TestOfferReader_processOffer(t *testing.T) {
	type fields struct {
		logger logger.LoggerInterface
		reader adapters.HandlerInterface
	}
	type args struct {
		header *RawHeader
		row    map[string]interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *RawOffer
		wantErr bool
	}{
		{
			name: "positive: row without required fields should be skipped",
			args: args{
				header: &RawHeader{
					Offer:    "Offers",
					Receiver: "Countries",
				},
				row: map[string]interface{}{
					"Offers":    nil,
					"Receiver":  "",
					"Countries": "US",
				},
			},
			want: nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &OfferReader{
				logger: tt.fields.logger,
				reader: tt.fields.reader,
			}
			got, err := o.processOffer(tt.args.header, tt.args.row)
			if (err != nil) != tt.wantErr {
				t.Errorf("processOffer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("processOffer() got = %v, want %v", got, tt.want)
			}
		})
	}
}
