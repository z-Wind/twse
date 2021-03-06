package twse

// StockInfo for Quotes
type StockInfo struct {
	// ServerResponse contains the HTTP response code and headers from the
	// server.
	ServerResponse `json:"-"`

	MilliSecond string `json:"tlong,omitempty"` //資料時間（毫秒）
	Date        string `json:"d,omitempty"`     //今日日期
	Time        string `json:"t,omitempty"`     //資料時間

	FullName string `json:"nf,omitempty"` //全名
	Name     string `json:"n,omitempty"`  //名字
	Symbol   string `json:"c,omitempty"`  //股要代碼
	Channel  string `json:"ch,omitempty"` //1101.tw

	BestBuyAmount      ListFloat64 `json:"f,omitempty"`         //最佳五檔賣出數量
	BestBuyPrice       ListFloat64 `json:"b,omitempty"`         //最佳五檔買入價格
	BestSellAmount     ListFloat64 `json:"g,omitempty"`         //最佳五檔買入數量
	BestSellPrice      ListFloat64 `json:"a,omitempty"`         //最佳五檔賣出價格
	TradePrice         Float64     `json:"z,omitempty,string"`  //最近成交價
	PreviousTradePrice Float64     `json:"pz,omitempty,string"` //前一個成交價
	YesterdayPrice     Float64     `json:"y,omitempty,string"`  //昨天收價
	Open               Float64     `json:"o,omitempty,string"`  //開盤價
	DayLow             Float64     `json:"l,omitempty,string"`  //今日最低
	DayHigh            Float64     `json:"h,omitempty,string"`  //今日最高

	Ex string `json:"ex,omitempty"` //上市上櫃
}
