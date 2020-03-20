package twse

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/pkg/errors"
)

// NewTimeseriesService get timeseries
func NewTimeseriesService(s *Service) *TimeseriesService {
	rs := &TimeseriesService{s: s}
	return rs
}

// TimeseriesService get timeseries
type TimeseriesService struct {
	s *Service
}

// MonthlyTWSE 上市股票查詢
// http://www.twse.com.tw/exchangeReport/STOCK_DAY?response=csv&date=20181230&stockNo=0050
// fix csv type
func (r *TimeseriesService) MonthlyTWSE(symbol string, date time.Time) *TimeseriesMonthlyTWSECall {
	c := &TimeseriesMonthlyTWSECall{
		DefaultCall: DefaultCall{
			s:         r.s,
			urlParams: url.Values{},
		},
	}

	c.urlParams.Set("response", "csv")
	c.urlParams.Set("date", date.Format("20060102"))
	c.urlParams.Set("stockNo", symbol)
	return c
}

// TimeseriesMonthlyTWSECall call function
type TimeseriesMonthlyTWSECall struct {
	DefaultCall

	symbol string
}

func (c *TimeseriesMonthlyTWSECall) doRequest() (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	reqHeaders.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	// 無需設定 http.Transport 已自帶，並自動解碼，若加上會產生亂碼
	// reqHeaders.Set("Accept-Encoding", "gzip, deflate, br")
	reqHeaders.Set("Accept-Language", "zh-TW,zh;q=0.9,en-US;q=0.8,en;q=0.7")

	var body io.Reader = nil
	urls := ResolveRelative(c.s.twseHost, "/exchangeReport/STOCK_DAY")
	urls += "?" + c.urlParams.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, errors.Wrapf(err, "http.NewRequest")
	}
	req.Header = reqHeaders

	return SendRequest(c.ctx, c.s.client, req)
}

// Do send request
func (c *TimeseriesMonthlyTWSECall) Do() (*TimeSeriesTWSEList, error) {
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

	ret := &TimeSeriesTWSEList{
		ServerResponse: ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := new([]*TimeSeriesTWSE)
	if err := DecodeResponseCSV(target, res); err != nil {
		return nil, errors.Wrapf(err, "DecodeResponseCSV")
	}

	ret.TimeSeries = *target

	return ret, nil
}

// MonthlyOTC 上櫃股票查詢
// http://www.tpex.org.tw/ch/stock/aftertrading/daily_trading_info/st43_download.php?d=108/02&stkno=4433&r=%d
// csv type
func (r *TimeseriesService) MonthlyOTC(symbol string, date time.Time) *TimeseriesMonthlyOTCCall {
	c := &TimeseriesMonthlyOTCCall{
		DefaultCall: DefaultCall{
			s:         r.s,
			urlParams: url.Values{},
		},
	}

	c.urlParams.Set("d", fmt.Sprintf("%d/%02d", date.Year()-1911, date.Month()))
	c.urlParams.Set("stkno", symbol)
	c.urlParams.Set("r", "%d")
	return c
}

// TimeseriesMonthlyOTCCall call function
type TimeseriesMonthlyOTCCall struct {
	DefaultCall

	symbol string
}

func (c *TimeseriesMonthlyOTCCall) doRequest() (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	reqHeaders.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	// 無需設定 http.Transport 已自帶，並自動解碼，若加上會產生亂碼
	// reqHeaders.Set("Accept-Encoding", "gzip, deflate, br")
	reqHeaders.Set("Accept-Language", "zh-TW,zh;q=0.9,en-US;q=0.8,en;q=0.7")

	var body io.Reader = nil
	urls := ResolveRelative(c.s.otcHost, "ch/stock/aftertrading/daily_trading_info/st43_download.php")
	urls += "?" + c.urlParams.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, errors.Wrapf(err, "http.NewRequest")
	}
	req.Header = reqHeaders

	return SendRequest(c.ctx, c.s.client, req)
}

// Do send request
func (c *TimeseriesMonthlyOTCCall) Do() (*TimeSeriesOTCList, error) {
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

	ret := &TimeSeriesOTCList{
		ServerResponse: ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := new([]*TimeSeriesOTC)
	if err := DecodeResponseCSV(target, res); err != nil {
		return nil, errors.Wrapf(err, "DecodeResponseCSV")
	}

	ret.TimeSeries = *target

	return ret, nil
}
