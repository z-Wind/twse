package twse

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"

	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/transform"
)

func TestTimeseriesMonthlyTWSECall_doRequest(t *testing.T) {
	client := clientTest("", http.StatusOK)
	twseTest, _ := New(client)

	tests := []struct {
		name    string
		c       *TimeseriesMonthlyTWSECall
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewTimeseriesService(twseTest).MonthlyTWSE("0050", time.Date(2020, time.March, 27, 0, 0, 0, 0, time.Local)), "https://www.twse.com.tw/exchangeReport/STOCK_DAY?date=20200327&response=csv&stockNo=0050", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rsp, err := tt.c.doRequest()
			if (err != nil) != tt.wantErr {
				t.Errorf("TimeseriesMonthlyTWSECall.doRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := rsp.Request.URL.String()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimeseriesMonthlyTWSECall.doRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestTimeseriesMonthlyTWSECall_Do(t *testing.T) {
	str, _, _ := transform.String(traditionalchinese.Big5.NewEncoder(), `"109年03月 0050 元大台灣50       各日成交資訊"
"日期","成交股數","成交金額","開盤價","最高價","最低價","收盤價","漲跌價差","成交筆數",
"109/03/02","27,894,641","2,437,814,508","87.50","88.25","86.85","87.35","-1.30","12,628",
"說明:"
"符號說明:+/-/X表示漲/跌/不比價"
"當日統計資訊含一般、零股、盤後定價、鉅額交易，不含拍賣、標購。"
"ETF證券代號第六碼為K、M、S、C者，表示該ETF以外幣交易。"`)
	client := clientTest(str, http.StatusOK)
	twseTest, _ := New(client)

	tests := []struct {
		name    string
		c       *TimeseriesMonthlyTWSECall
		want    []*TimeSeriesTWSE
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewTimeseriesService(twseTest).MonthlyTWSE("0050", time.Now()), []*TimeSeriesTWSE{
			&TimeSeriesTWSE{Time: Time(time.Date(2020, 3, 2, 0, 0, 0, 0, time.UTC)), Volume: 2.7894641e+07, Turnover: 2.437814508e+09, Open: 87.5, High: 88.25, Low: 86.85, Close: 87.35, Change: -1.3, Transactions: 12628},
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rsp, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("TimeseriesMonthlyTWSECall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := rsp.TimeSeries
			fmt.Printf("%+v\n%+v\n", got[0], tt.want[0])
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimeseriesMonthlyTWSECall.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeseriesMonthlyOTCCall_doRequest(t *testing.T) {
	client := clientTest("", http.StatusOK)
	twseTest, _ := New(client)

	tests := []struct {
		name    string
		c       *TimeseriesMonthlyOTCCall
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewTimeseriesService(twseTest).MonthlyOTC("5483", time.Date(2020, time.March, 27, 0, 0, 0, 0, time.Local)), "https://www.tpex.org.tw/ch/stock/aftertrading/daily_trading_info/st43_download.php?d=109%2F03&r=%25d&stkno=5483", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rsp, err := tt.c.doRequest()
			if (err != nil) != tt.wantErr {
				t.Errorf("TimeseriesMonthlyOTCCall.doRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := rsp.Request.URL.String()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimeseriesMonthlyOTCCall.doRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeseriesMonthlyOTCCall_Do(t *testing.T) {
	str, _, _ := transform.String(traditionalchinese.Big5.NewEncoder(), `個股日成交資訊  
股票代號:5483
股票名稱:中美晶
資料日期:109/03
日 期,成交仟股,成交仟元,開盤,最高,最低,收盤,漲跌,筆數
"109/03/02","13,630","1,443,826","102.00","108.50","101.50","108.50","3.50","6,796"
共19筆`)
	client := clientTest(str, http.StatusOK)
	twseTest, _ := New(client)

	tests := []struct {
		name    string
		c       *TimeseriesMonthlyOTCCall
		want    []*TimeSeriesOTC
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewTimeseriesService(twseTest).MonthlyOTC("5483", time.Now()), []*TimeSeriesOTC{
			&TimeSeriesOTC{Time: Time(time.Date(2020, 3, 2, 0, 0, 0, 0, time.UTC)), Volume: 13630, Turnover: 1.443826e+06, Open: 102, High: 108.5, Low: 101.5, Close: 108.5, Change: 3.5, Transactions: 6796},
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rsp, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("TimeseriesMonthlyOTCCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := rsp.TimeSeries
			fmt.Printf("%+v\n%+v\n", got[0], tt.want[0])
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimeseriesMonthlyOTCCall.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}
