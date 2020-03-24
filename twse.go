package twse

import (
	"errors"
	"log"
	"net"
	"net/http"
	"net/http/cookiejar"
	"time"
)

// const strings
const (
	// UserAgent is the header string used to identify this package.
	userAgent = `Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:62.0) Gecko/20100101 Firefox/62.0`

	TWSEURL  = "https://mis.twse.com.tw"
	TWSEHOST = "https://www.twse.com.tw"
	OTCHOST  = "https://www.tpex.org.tw"
)

// Service TWSE api
type Service struct {
	client *http.Client

	twseURL  string // API endpoint base URL
	twseHost string // API endpoint base URL
	otcHost  string // API endpoint base URL

	Quotes     *QuotesService
	Timeseries *TimeseriesService
}

// GetClient get client
func GetClient() *http.Client {
	cookieJar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			Dial: (&net.Dialer{
				Timeout:   0,
				KeepAlive: 0,
			}).Dial,
			TLSHandshakeTimeout: 10 * time.Second,
		},
		Jar: cookieJar,
	}

	// 前置作業，取得 cookies
	urls := ResolveRelative(TWSEURL, "stock/index.jsp")
	_, err = client.Get(urls)
	if err != nil {
		log.Fatalf("c.s.client.Get(%s): %s", urls, err)
	}

	return client
}

// New TWSE API server
func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, twseURL: TWSEURL, twseHost: TWSEHOST, otcHost: OTCHOST}
	s.Quotes = NewQuotesService(s)
	s.Timeseries = NewTimeseriesService(s)

	return s, nil
}

func (s *Service) userAgent() string {
	return userAgent
}
