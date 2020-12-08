package twse

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// ServerResponse is embedded in each Do response and
// provides the HTTP status code and header sent by the server.
type ServerResponse struct {
	// HTTPStatusCode is the server's response status code. When using a
	// resource method's Do call, this will always be in the 2xx range.
	HTTPStatusCode int
	// Header contains the response header fields from the server.
	Header http.Header
}

// DefaultCall DefaultCall function
type DefaultCall struct {
	s         *Service
	urlParams url.Values
	ctx       context.Context
	header    http.Header
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *DefaultCall) Context(ctx context.Context) *DefaultCall {
	c.ctx = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *DefaultCall) Header() http.Header {
	if c.header == nil {
		c.header = make(http.Header)
	}
	return c.header
}

// Float64 unmarshal string to Float64
type Float64 float64

func (f *Float64) unmarshal(data []byte) error {
	s := string(data)
	s = strings.ReplaceAll(s, ",", "")
	s = strings.ReplaceAll(s, " ", "")
	if s == "--" || s == `-` {
		*f = 0
		return nil
	}

	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return errors.Wrapf(err, "strconv.ParseFloat %s", s)
	}

	*f = Float64(v)

	return nil
}

// UnmarshalJSON process Date
func (f *Float64) UnmarshalJSON(data []byte) error {
	return f.unmarshal(data)
}

// UnmarshalCSV process Date
func (f *Float64) UnmarshalCSV(data []byte) error {
	return f.unmarshal(data)
}

// ListFloat64 unmarshal string to ListInt
type ListFloat64 []float64

// UnmarshalJSON process Date
func (l *ListFloat64) UnmarshalJSON(data []byte) error {
	result := []float64{}
	str := string(data)
	str = strings.ReplaceAll(str, `"`, "")
	sList := strings.Split(str, "_")

	for _, s := range sList[:len(sList)-1] {
		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return errors.Wrapf(err, "strconv.ParseFloat %s in %v(%s)", s, sList, str)
		}
		result = append(result, f)
	}

	*l = result

	return nil
}
