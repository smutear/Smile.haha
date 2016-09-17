package code

import (
	"compress/gzip"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// WebPage wraps necessary data to make a POST or GET http request,
// and reads its reponse. It provides additional functionalities to
// assist parsing the results.
type WebPage struct {
	// request
	Url      string
	Method   string // "POST" or "GET" supported
	Proxy    string
	User     string // support for basic auth (not encrypted)
	Password string
	// may set compression and other options here
	//	Header = map[string][]string{
	//		"Accept-Encoding": {"gzip, deflate"},
	//		"Accept-Language": {"en-us"},
	//		"Connection": {"keep-alive"},
	//	}
	Header   http.Header
	BodyType string        // e.g. "text/xml", "application/json", "application/x-www-form-urlencoded"
	Body     io.Reader     // the data sent when using POST method
	Timeout  time.Duration // time out

	// response
	StatusCode    int    // response status code, client should check it to decide the next move
	Status        string // the status message returned
	Cookies       []*http.Cookie
	RespHeader    http.Header
	RespBody      []byte // try to read even if stauts code is not 200, since some web server returns error message which might be useful
	CheckRedirect func(req *http.Request, via []*http.Request) error
}

// DoRequest wraps the common http request paradigm to get the request result.
func (w *WebPage) DoRequest() error {
	// no need to check parameters, NewRequest would do it.
	req, err := http.NewRequest(w.Method, w.Url, w.Body)
	if err != nil {
		return err
	}
	if w.Header != nil {
		req.Header = w.Header
	}
	if w.User != "" && w.Password != "" {
		req.SetBasicAuth(w.User, w.Password)
	}
	// add body type
	if w.BodyType != "" {
		req.Header.Set("Content-Type", w.BodyType)
	}
	client := http.DefaultClient
	if w.Proxy != "" {
		proxy, _ := url.Parse(w.Proxy)
		client = &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxy)}}
	}

	// enforce the CheckRedirect func
	client.CheckRedirect = w.CheckRedirect

	if w.Timeout > 0 {
		client.Timeout = w.Timeout
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	w.Cookies = resp.Cookies()
	w.RespHeader = resp.Header
	w.StatusCode, w.Status = resp.StatusCode, resp.Status
	var page []byte
	if enc := resp.Header.Get("Content-Encoding"); enc == "gzip" {
		reader, err := gzip.NewReader(resp.Body)
		if err != nil {
			return err
		}
		page, err = ioutil.ReadAll(reader)
	} else {
		page, err = ioutil.ReadAll(resp.Body)
	}

	if err != nil {
		return err
	}
	w.RespBody = page
	return nil
}

// could overwrite existing headers
func (w *WebPage) SetHeader(key string, value string) {
	w.Header.Set(key, value)
}
