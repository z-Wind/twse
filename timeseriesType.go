package twse

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
)

// Time redefine time.time for Unmarshal
type Time time.Time

// UnmarshalCSV process Date
func (t *Time) UnmarshalCSV(data []byte) error {
	var year, month, day int
	_, err := fmt.Sscanf(string(data), "%d/%d/%d", &year, &month, &day)
	if err != nil {
		return errors.Wrapf(err, "fmt.Sscanf %s", string(data))
	}

	// 民國轉西元
	s := fmt.Sprintf("%d-%02d-%02d", year+1911, month, day)

	// timeSeriesDateFormats are the expected date formats in time series data
	timeSeriesDateFormats := []string{
		"2006-01-02",
	}

	d, err := parseDate(s, timeSeriesDateFormats...)
	if err != nil {
		return errors.Wrapf(err, "error parsing timestamp %s", s)
	}
	*t = Time(d)

	return nil
}

// TimeSeriesTWSE time series
type TimeSeriesTWSE struct {
	Time         Time    `csv:"日期"`
	Volume       Float64 `csv:"成交股數"`
	Turnover     Float64 `csv:"成交金額"`
	Open         Float64 `csv:"開盤價"`
	High         Float64 `csv:"最高價"`
	Low          Float64 `csv:"最低價"`
	Close        Float64 `csv:"收盤價"`
	Change       Float64 `csv:"漲跌價差"`
	Transactions Float64 `csv:"成交筆數"`
}

// TimeSeriesTWSEList TimeSeries List
type TimeSeriesTWSEList struct {
	// ServerResponse contains the HTTP response code and headers from the
	// server.
	ServerResponse `csv:"-"`

	TimeSeries []*TimeSeriesTWSE
}

// TimeSeriesOTC time series
type TimeSeriesOTC struct {
	Time         Time    `csv:"日 期"`
	Volume       Float64 `csv:"成交仟股"`
	Turnover     Float64 `csv:"成交仟元"`
	Open         Float64 `csv:"開盤"`
	High         Float64 `csv:"最高"`
	Low          Float64 `csv:"最低"`
	Close        Float64 `csv:"收盤"`
	Change       Float64 `csv:"漲跌"`
	Transactions Float64 `csv:"筆數"`
}

// TimeSeriesOTCList TimeSeries List
type TimeSeriesOTCList struct {
	// ServerResponse contains the HTTP response code and headers from the
	// server.
	ServerResponse `csv:"-"`

	TimeSeries []*TimeSeriesOTC
}
