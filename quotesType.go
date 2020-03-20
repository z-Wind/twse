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

	BestBuyAmount      string  `json:"f,omitempty"`         //最佳五檔賣出數量
	BestBuyPrice       string  `json:"b,omitempty"`         //最佳五檔買入價格
	BestSellAmount     string  `json:"g,omitempty"`         //最佳五檔買入數量
	BestSellPrice      string  `json:"a,omitempty"`         //最佳五檔賣出價格
	TradePrice         float64 `json:"z,omitempty,string"`  //最近成交價
	PreviousTradePrice float64 `json:"pz,omitempty,string"` //前一個成交價
	YesterdayPrice     float64 `json:"y,omitempty,string"`  //昨天收價
	Open               float64 `json:"o,omitempty,string"`  //開盤價
	DayLow             float64 `json:"l,omitempty,string"`  //今日最低
	DayHigh            float64 `json:"h,omitempty,string"`  //今日最高

	Ex string `json:"ex,omitempty"` //上市上櫃
}
