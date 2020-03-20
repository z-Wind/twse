package twse

import (
	"reflect"
	"testing"
	"time"
)

func TestTimeseriesMonthlyTWSECall_Do(t *testing.T) {
	tests := []struct {
		name    string
		c       *TimeseriesMonthlyTWSECall
		want    *TimeSeriesTWSEList
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewTimeseriesService(twse).MonthlyTWSE("0050", time.Now()), &TimeSeriesTWSEList{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("TimeseriesMonthlyTWSECall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimeseriesMonthlyTWSECall.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeseriesMonthlyOTCCall_Do(t *testing.T) {
	tests := []struct {
		name    string
		c       *TimeseriesMonthlyOTCCall
		want    *TimeSeriesOTCList
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewTimeseriesService(twse).MonthlyOTC("4433", time.Now()), &TimeSeriesOTCList{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("TimeseriesMonthlyOTCCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimeseriesMonthlyOTCCall.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}
