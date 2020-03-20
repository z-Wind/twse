package twse

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/pkg/errors"
)

// NewQuotesService get last price
// http://mis.twse.com.tw/stock/api/getStockInfo.jsp?ex_ch=tse_1101.tw&json=1&delay=0&_=1539865363091
func NewQuotesService(s *Service) *QuotesService {
	rs := &QuotesService{s: s}
	return rs
}

// QuotesService get last price
// http://mis.twse.com.tw/stock/api/getStockInfo.jsp?ex_ch=tse_1101.tw&json=1&delay=0&_=1539865363091
type QuotesService struct {
	s *Service
}

// GetStockInfoTWSE 上市股票查詢
// http://mis.twse.com.tw/stock/api/getStockInfo.jsp?ex_ch=tse_1101.tw&json=1&delay=0&_=1539865363091
func (r *QuotesService) GetStockInfoTWSE(symbol string) *QuotesGetStockInfoCall {
	return r.getStockInfo("tse", symbol)
}

// GetStockInfoOTC 上櫃股票查詢
// http://mis.twse.com.tw/stock/api/getStockInfo.jsp?ex_ch=tse_1101.tw&json=1&delay=0&_=1539865363091
func (r *QuotesService) GetStockInfoOTC(symbol string) *QuotesGetStockInfoCall {
	return r.getStockInfo("otc", symbol)
}

// http://mis.twse.com.tw/stock/api/getStockInfo.jsp?ex_ch=tse_1101.tw&json=1&delay=0&_=1539865363091
func (r *QuotesService) getStockInfo(market, symbol string) *QuotesGetStockInfoCall {
	c := &QuotesGetStockInfoCall{
		DefaultCall: DefaultCall{
			s:         r.s,
			urlParams: url.Values{},
		},

		symbol: symbol,
	}

	millis := time.Now().UnixNano() / int64(time.Millisecond)

	c.urlParams.Set("ex_ch", fmt.Sprintf("%s_%s.tw", market, symbol))
	c.urlParams.Set("json", "1")
	c.urlParams.Set("delay", "0")
	c.urlParams.Set("_", fmt.Sprintf("%d", millis))
	return c
}

// QuotesGetStockInfoCall call function
type QuotesGetStockInfoCall struct {
	DefaultCall

	symbol string
}

func (c *QuotesGetStockInfoCall) doRequest() (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	reqHeaders.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	// 無需設定 http.Transport 已自帶，並自動解碼，若加上會產生亂碼
	// reqHeaders.Set("Accept-Encoding", "gzip, deflate, br")
	reqHeaders.Set("Accept-Language", "zh-TW,zh;q=0.9,en-US;q=0.8,en;q=0.7")
	reqHeaders.Set("Referer", "http://mis.twse.com.tw/stock/fibest.jsp?stock="+c.symbol)
	reqHeaders.Set("X-Requested-With", "XMLHttpRequest")

	var body io.Reader = nil
	urls := ResolveRelative(c.s.twseURL, "stock/api/getStockInfo.jsp")
	urls += "?" + c.urlParams.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, errors.Wrapf(err, "http.NewRequest")
	}
	req.Header = reqHeaders

	return SendRequest(c.ctx, c.s.client, req)
}

// Do send request
func (c *QuotesGetStockInfoCall) Do() (*StockInfo, error) {
	res, err := c.doRequest()
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, errors.Wrapf(err, "doRequest")
	}
	defer res.Body.Close()
	if err := CheckResponse(res); err != nil {
		return nil, errors.Wrapf(err, "CheckResponse")
	}

	target := new(struct {
		StockInfoList []*StockInfo `json:"msgArray"`
	})
	if err := DecodeResponseJSON(target, res); err != nil {
		return nil, errors.Wrapf(err, "DecodeResponseJSON")
	}

	ret := (*target).StockInfoList[0]
	ret.ServerResponse = ServerResponse{
		Header:         res.Header,
		HTTPStatusCode: res.StatusCode,
	}

	return ret, nil
}
