package twse

import (
	"reflect"
	"testing"
)

func TestQuotesGetStockInfoCall_Do(t *testing.T) {
	tests := []struct {
		name    string
		c       *QuotesGetStockInfoCall
		want    *StockInfo
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewQuotesService(twse).GetStockInfoTWSE("0050"), &StockInfo{}, false},
        {"Test", NewQuotesService(twse).GetStockInfoTWSE("VTI"), nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("QuotesGetStockInfoCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuotesGetStockInfoCall.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}
