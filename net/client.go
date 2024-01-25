package net

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/cookiejar"
	"strings"
	"time"
	. "wu"
)

type HttpClient struct {
	Client         *http.Client
	Cookies        []*http.Cookie
	HttpCookieJar  *cookiejar.Jar
	DefaultCookies string
	// Transport *http.Transport
}

const (
	USER_AGENT              = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_5) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.1.1 Safari/605.1.15"
	MaxIdleConns        int = 1024
	MaxIdleConnsPerHost int = 256
	RequestTimeout      int = 160
	IdleConnTimeout     int = 60
)

var (
	myCookieJar, _ = cookiejar.New(nil)
)

func NewHttpClient() *HttpClient {
	a := &HttpClient{}
	a.DefaultCookies = ``
	a.Cookies = nil
	a.Client = createHTTPClient()
	return a
}

// createHTTPClient for connection re-use
func createHTTPClient() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			Proxy:               http.ProxyFromEnvironment,
			DialContext:         (&net.Dialer{Timeout: time.Duration(RequestTimeout) * time.Second, KeepAlive: time.Duration(RequestTimeout) * time.Second}).DialContext,
			MaxIdleConns:        MaxIdleConns,
			MaxIdleConnsPerHost: MaxIdleConnsPerHost,
			IdleConnTimeout:     time.Duration(IdleConnTimeout) * time.Second,
		},
		Timeout: time.Duration(RequestTimeout) * time.Second,
		Jar:     myCookieJar,
	}
	return client
}

// set common http headers
func (a *HttpClient) SetHeaderCommon(req *http.Request) {
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.8")
	req.Header.Set("User-Agent", USER_AGENT)
	req.Header.Set("Connection", "keep-alive")
}

func (a *HttpClient) SetHeaderAcceptAll(req *http.Request) {
	req.Header.Set("Accept", "*/*")
}
func (a *HttpClient) SetHeaderAcceptEncoding(req *http.Request) {
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
}
func (a *HttpClient) SetHeaderAcceptLanguage(req *http.Request) {
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.8")
}

// set conn keep alive
func (a *HttpClient) SetHeaderConnKeepAlive(req *http.Request) {
	req.Header.Set("Connection", "keep-alive")
}

// set conn close
func (a *HttpClient) SetHeaderConnClose(req *http.Request) {
	req.Header.Set("Connection", "close")
}

// set post http header
func (a *HttpClient) SetHeaderPOST(req *http.Request) {
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
}

// set XMLHttpRequest http header
func (a *HttpClient) SetHeaderXMLHttpRequest(req *http.Request) {
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
}

// set user agent
func (a *HttpClient) SetHeaderUserAgent(req *http.Request, userAgent string) {
	if userAgent == "" {
		userAgent = USER_AGENT
	}
	req.Header.Set("User-Agent", userAgent)
}

// set cookies
func (a *HttpClient) SetHeaderCookies(req *http.Request, cookies string) {
	req.Header.Set("Cookie", cookies)
}

// set Referer
func (a *HttpClient) SetHeaderReferer(req *http.Request, referer string) {
	req.Header.Set("Referer", referer)
}

// set Origin
func (a *HttpClient) SetHeaderOrigin(req *http.Request, origin string) {
	req.Header.Set("Origin", origin)
}

func (a *HttpClient) GET(url string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if Err(err) {
		return "", err
	}
	a.SetHeaderCommon(req)
	if a.DefaultCookies != "" {
		a.SetHeaderCookies(req, a.DefaultCookies)
	}
	resp, err := a.Client.Do(req)
	if Err(err) {
		println("GET a.Client.Do err", url)
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if Err(err) {
		println("GET ioutil.ReadAll err", url)
		return "", err
	}
	a.Cookies = myCookieJar.Cookies(req.URL)
	return string(body), err
}

func (a *HttpClient) POST(url, postData string) (string, error) {
	req, err := http.NewRequest("POST", url, strings.NewReader(postData))
	if Err(err) {
		return "", err
	}
	a.SetHeaderCommon(req)
	a.SetHeaderPOST(req)
	if a.DefaultCookies != "" {
		a.SetHeaderCookies(req, a.DefaultCookies)
	}
	resp, err := a.Client.Do(req)
	if Err(err) {
		println("POST a.Client.Do err", url)
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if Err(err) {
		println("POST ioutil.ReadAll err", url)
		return "", err
	}
	a.Cookies = myCookieJar.Cookies(req.URL)
	return string(body), err
}

// post for xzsec
func (a *HttpClient) XZSEC_POST(url, postData string) (string, error) {
	req, err := http.NewRequest("POST", url, strings.NewReader(postData))
	if Err(err) {
		return "", err
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "zh-CN;q=0.8,zh;q=0.7")
	req.Header.Set("User-Agent", USER_AGENT)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Origin", "https://jy.xzsec.com")
	req.Header.Set("Referer", "https://jy.xzsec.com")
	req.Header.Set("Connection", "keep-alive")

	if a.DefaultCookies != "" {
		a.SetHeaderCookies(req, a.DefaultCookies)
	}
	resp, err := a.Client.Do(req)
	if Err(err) {
		// println("POST a.Client.Do err", url)
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if Err(err) {
		// println("POST ioutil.ReadAll err", url)
		return "", err
	}
	a.Cookies = myCookieJar.Cookies(req.URL)
	return string(body), err
}

// get all cookies string
func (a *HttpClient) GetAllCookiesString() string {
	cookieStr := ""
	cookieNum := len(a.Cookies)
	if cookieNum == 0 {
		return ""
	}
	// Printfln("cookieNum = %d", cookieNum)
	for i := 0; i < cookieNum; i++ {
		var ck *http.Cookie = a.Cookies[i]
		if i != cookieNum {
			cookieStr = fmt.Sprintf("%v%s=%v; ", cookieStr, ck.Name, ck.Value)
		} else {
			// last cookie
			cookieStr = fmt.Sprintf("%v%s=%v", cookieStr, ck.Name, ck.Value)
		}
	}
	return cookieStr
}

func (a *HttpClient) PrintCookies() {
	cookieNum := len(a.Cookies)
	Printfln("cookieNum = %d", cookieNum)
	for i := 0; i < cookieNum; i++ {
		var ck *http.Cookie = a.Cookies[i]
		Printfln("\n------ Cookie [%d]------", i)
		Printfln("Name\t\t= %s", ck.Name)
		Printfln("Value\t\t= %v", ck.Value)
		Printfln("Path\t\t= %s", ck.Path)
		Printfln("Domain\t\t= %s", ck.Domain)
		Printfln("Expires\t\t= %s", ck.Expires)
		Printfln("MaxAge\t\t= %d", ck.MaxAge)
		Printfln("Secure\t\t= %t", ck.Secure)
		Printfln("HttpOnly\t= %t", ck.HttpOnly)
		Printfln("Raw\t\t= %s", ck.Raw)
		Printfln("RawExpires\t= %s", ck.RawExpires)
		Printfln("Unparsed\t= %s", ck.Unparsed)
	}
}
