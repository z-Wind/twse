package twse

import (
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	twseReal *Service
)

func init() {
	client := GetClient()

	var err error
	twseReal, err = New(client)
	if err != nil {
		panic(err)
	}
}

type TestTransport struct {
	body       string
	statusCode int
}

// RoundTrip add apikey
func (t *TestTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	var res http.Response
	res.StatusCode = t.statusCode
	res.Body = ioutil.NopCloser(strings.NewReader(t.body))
	res.Header = http.Header{}
	res.Request = req

	return &res, nil
}

func clientTest(body string, statuscode int) *http.Client {
	transport := &TestTransport{body: body, statusCode: statuscode}

	client := &http.Client{
		Transport: transport,
	}

	return client
}
