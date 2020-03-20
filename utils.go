package twse

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os/exec"
	"path"
	"regexp"
	"runtime"
	"time"

	"github.com/jszwec/csvutil"
	"github.com/pkg/errors"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/transform"
)

// openURL opens a browser window to the specified location.
// This code originally appeared at:
//   http://stackoverflow.com/questions/10377243/how-can-i-launch-a-process-that-is-not-a-file-in-go
func openURL(url string) error {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("Cannot open URL %s on this platform", url)
	}
	return err
}

// CheckResponse returns an error (of type *Error) if the response
// status code is not 2xx.
func CheckResponse(res *http.Response) error {
	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		return nil
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return errors.Wrapf(err, "ioutil.ReadAll")
	}

	errReply := &errorReply{}
	err = json.Unmarshal(b, errReply)
	if err != nil {
		return nil
	}

	return &Error{
		Code:    res.StatusCode,
		Message: errReply.Message,
		Body:    string(b),
		Header:  res.Header,
	}
}

// DecodeResponseJSON decodes the body of res into target. If there is no body,
// target is unchanged.
func DecodeResponseJSON(target interface{}, res *http.Response) error {
	if res.StatusCode == http.StatusNoContent {
		return nil
	}
	// for test
	// b, _ := ioutil.ReadAll(res.Body)
	// ioutil.WriteFile("tt.txt", b, 0666)
	//fmt.Printf("%s\n", string(b))
	//fmt.Printf("====================================\n")
	//return json.NewDecoder(bytes.NewBuffer(b)).Decode(target)

	return json.NewDecoder(res.Body).Decode(target)
}

// DecodeResponseCSV decodes the body of res into target. If there is no body,
// target is unchanged.
func DecodeResponseCSV(target interface{}, res *http.Response) error {
	if res.StatusCode == http.StatusNoContent {
		return nil
	}

	O := transform.NewReader(res.Body, traditionalchinese.Big5.NewDecoder())

	b, err := ioutil.ReadAll(O)
	if err != nil {
		return errors.Wrapf(err, "ioutil.ReadAll")
	}

	s := string(b)
	// 上市處理
	s = regexp.MustCompile(`".*各日成交資訊"`).ReplaceAllString(s, "")
	// 移除上市最後的說明
	s = regexp.MustCompile(`"[^,0-9]{40,}"`).ReplaceAllString(s, "")

	// 上櫃處理
	s = regexp.MustCompile(`共\d+筆`).ReplaceAllString(s, "")
	// 移除上櫃開頭說明
	s = regexp.MustCompile(`個股日成交資訊[^-]*日期.*`).ReplaceAllString(s, "")

	// 移除空白行
	s = regexp.MustCompile(`^\s+`).ReplaceAllString(s, "")

	// lines := strings.Split(s, "\n")
	// // 移除第一行 Header
	// lines = lines[1:]
	// s = strings.Join(lines, "\n")

	fmt.Printf("%s\n", s)
	fmt.Printf("====================================\n")

	return csvutil.Unmarshal([]byte(s), target)
}

// SendRequest sends a single HTTP request using the given client.
// If ctx is non-nil, it sends the request with req.WithContext
func SendRequest(ctx context.Context, client *http.Client, req *http.Request) (*http.Response, error) {
	if ctx == nil {
		return client.Do(req)
	}

	resp, err := client.Do(req.WithContext(ctx))
	// If we got an error, and the context has been canceled,
	// the context's error is probably more useful.
	if err != nil {
		select {
		case <-ctx.Done():
			err = ctx.Err()
		default:
		}
	}

	return resp, err
}

// JSONReader convert struct to reader for json request
func JSONReader(v interface{}) (io.Reader, error) {
	buf := new(bytes.Buffer)

	err := json.NewEncoder(buf).Encode(v)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

// ResolveRelative join path
func ResolveRelative(basePath string, elem ...string) string {
	u, err := url.Parse(basePath)
	if err != nil {
		panic(fmt.Sprintf("url.Parse failed to parse %q", basePath))
	}

	u.Path = path.Join(u.Path, path.Join(elem...))

	return u.String()
}

// parseDate parses a date value from a string.
// An error is returned if the value is not in one of the dateFormat formats.
func parseDate(v string, dateFormat ...string) (time.Time, error) {
	for _, format := range dateFormat {
		t, err := time.Parse(format, v)
		if err != nil {
			continue
		}
		return t, nil
	}
	return time.Time{}, errors.Errorf("applicable date format not found for date %s", v)
}
