package twse

import (
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"testing"
)

func Test_Quotes(t *testing.T) {
	server := NewQuotesService(twseReal)

	GetStockInfoTWSECall := server.GetStockInfoTWSE("0050")
	got, err := GetStockInfoTWSECall.Do()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v", got)

	GetStockInfoTWSECall = server.GetStockInfoTWSE("006208")
	got, err = GetStockInfoTWSECall.Do()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v", got)

	GetStockInfoOTCCall := server.GetStockInfoOTC("5483")
	got, err = GetStockInfoOTCCall.Do()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v", got)
}

func TestQuotesGetStockInfoCall_doRequest(t *testing.T) {
	client := clientTest("", http.StatusOK)
	twseTest, _ := New(client)

	tests := []struct {
		name    string
		c       *QuotesGetStockInfoCall
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"OTC", NewQuotesService(twseTest).GetStockInfoOTC("5483"), "https://mis.twse.com.tw/stock/api/getStockInfo.jsp?_=1585269276147&delay=0&ex_ch=otc_5483.tw&json=1", false},
		{"TWSE", NewQuotesService(twseTest).GetStockInfoTWSE("0050"), "https://mis.twse.com.tw/stock/api/getStockInfo.jsp?_=1585269276147&delay=0&ex_ch=tse_0050.tw&json=1", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rsp, err := tt.c.doRequest()
			if (err != nil) != tt.wantErr {
				t.Errorf("QuotesGetStockInfoCall.doRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := rsp.Request.URL.String()
			// 將時間取代以固定之
			got = regexp.MustCompile(`_=\d+&`).ReplaceAllString(got, "_=1585269276147&")
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuotesGetStockInfoCall.doRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuotesGetStockInfoCall_Do(t *testing.T) {
	client := clientTest(`{"queryTime":{"stockInfoItem":1391,"sessionKey":"tse_0050.tw_20200327|","sessionStr":"UserSession","sysDate":"20200327","sessionFromTime":-1,"stockInfo":123783,"showChart":false,"sessionLatestTime":-1,"sysTime":"08:34:46"},"referer":"","rtmessage":"OK","exKey":"if_tse_0050.tw_zh-tw.null","msgArray":[{"n":"元大台灣50","g":"7_5_8_2_13_","u":"84.9000","mt":"128345","o":"-","ps":"251","a":"78.0500_78.1000_78.1500_78.2000_78.2500_","tlong":"1585269282000","t":"08:34:42","it":"02","ch":"0050.tw","b":"78.0000_77.9500_77.9000_77.8500_77.8000_","f":"1_5_1_14_4_","w":"69.5000","pz":"78.0000","l":"-","c":"0050","v":"0","d":"20200327","tv":"-","tk1":"0050.tw_tse_20200327_B_9999999997","ts":"1","nu":"http://www.yuantaetfs.com/#/RtNav/Index","nf":"元大台灣卓越50證券投資信託基金","y":"77.2000","p":"0","ip":"0","z":"-","s":"-","h":"-","ex":"tse"}],"userDelay":5000,"rtcode":"0000","cachedAlive":96622}`, http.StatusOK)
	twseTest, _ := New(client)

	tests := []struct {
		name    string
		c       *QuotesGetStockInfoCall
		want    *StockInfo
		wantErr bool
	}{
		// TODO: Add test cases.
		{"TWSE", NewQuotesService(twseTest).GetStockInfoTWSE("0050"), &StockInfo{
			ServerResponse:     ServerResponse{HTTPStatusCode: 200, Header: http.Header{}},
			MilliSecond:        "1585269282000",
			Date:               "20200327",
			Time:               "08:34:42",
			FullName:           "元大台灣卓越50證券投資信託基金",
			Name:               "元大台灣50",
			Symbol:             "0050",
			Channel:            "0050.tw",
			BestBuyAmount:      []float64{1, 5, 1, 14, 4},
			BestBuyPrice:       []float64{78.0000, 77.9500, 77.9000, 77.8500, 77.8000},
			BestSellAmount:     []float64{7, 5, 8, 2, 13},
			BestSellPrice:      []float64{78.0500, 78.1000, 78.1500, 78.2000, 78.2500},
			TradePrice:         0,
			PreviousTradePrice: 78,
			YesterdayPrice:     77.2,
			Open:               0,
			DayLow:             0,
			DayHigh:            0,
			Ex:                 "tse",
		}, false},
		{"Symbol Not Found", NewQuotesService(twseReal).GetStockInfoTWSE("VTI"), nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("QuotesGetStockInfoCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuotesGetStockInfoCall.Do() = \n%+v\n, want \n%+v", got, tt.want)
			}
		})
	}

	client = clientTest(`{"queryTime":{"stockInfoItem":1620,"sessionKey":"otc_5483.tw_20200327|","sessionStr":"UserSession","sysDate":"20200327","sessionFromTime":-1,"stockInfo":128464,"showChart":false,"sessionLatestTime":-1,"sysTime":"08:36:36"},"referer":"","rtmessage":"OK","exKey":"if_otc_5483.tw_zh-tw.null","msgArray":[{"n":"中美晶","g":"87_1_97_5_1_","u":"88.3000","mt":"329248","o":"-","ps":"174","a":"83.7000_83.8000_83.9000_84.0000_84.1000_","tlong":"1585269391000","t":"08:36:31","it":"12","ch":"5483.tw","b":"83.5000_83.1000_83.0000_82.0000_81.9000_","f":"2_9_9_38_3_","w":"72.3000","pz":"83.5000","l":"-","c":"5483","v":"0","d":"20200327","tv":"-","tk1":"5483.tw_otc_20200327_B_9999996079","ts":"1","nf":"中美矽晶製品股份有限公司","y":"80.3000","p":"0","i":"24","ip":"0","z":"-","s":"-","h":"-","ex":"otc"}],"userDelay":5000,"rtcode":"0000","cachedAlive":32382}`, http.StatusOK)
	twseTest, _ = New(client)

	tests = []struct {
		name    string
		c       *QuotesGetStockInfoCall
		want    *StockInfo
		wantErr bool
	}{
		// TODO: Add test cases.

		{"OTC", NewQuotesService(twseTest).GetStockInfoOTC("5483"), &StockInfo{
			ServerResponse:     ServerResponse{HTTPStatusCode: 200, Header: http.Header{}},
			MilliSecond:        "1585269391000",
			Date:               "20200327",
			Time:               "08:36:31",
			FullName:           "中美矽晶製品股份有限公司",
			Name:               "中美晶",
			Symbol:             "5483",
			Channel:            "5483.tw",
			BestBuyAmount:      []float64{2, 9, 9, 38, 3},
			BestBuyPrice:       []float64{83.5, 83.1, 83, 82, 81.9},
			BestSellAmount:     []float64{87, 1, 97, 5, 1},
			BestSellPrice:      []float64{83.7, 83.8, 83.9, 84, 84.1},
			TradePrice:         0,
			PreviousTradePrice: 83.5,
			YesterdayPrice:     80.3,
			Open:               0,
			DayLow:             0,
			DayHigh:            0,
			Ex:                 "otc",
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("QuotesGetStockInfoCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuotesGetStockInfoCall.Do() = \n%+v\n, want \n%+v", got, tt.want)
			}
		})
	}
}
