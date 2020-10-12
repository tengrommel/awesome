package main

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var addr = "127.0.0.1:2002"

func main() {
	rs1 := "http://127.0.0.1:2003"
	url1, err1 := url.Parse(rs1)
	if err1 != nil {
		log.Println(err1)
	}
	rs2 := "http://127.0.0.1:2004"
	url2, err2 := url.Parse(rs2)
	if err2 != nil {
		log.Println(err2)
	}
	urls := []*url.URL{url1, url2}
	proxy := NewMultipleHostsReverseProxy(urls)
	log.Println("Starting httpserver at " + addr)
	log.Fatal(http.ListenAndServe(addr, proxy))
}

var transport = &http.Transport{
	DialContext: (&net.Dialer{
		Timeout:   30 * time.Second, //连接超时
		KeepAlive: 30 * time.Second, //长连接超时时间
	}).DialContext,
	MaxIdleConns:          100,              //最大空闲连接
	IdleConnTimeout:       90 * time.Second, //空闲超时时间
	TLSHandshakeTimeout:   10 * time.Second, //tls握手超时时间
	ExpectContinueTimeout: 1 * time.Second,  //100-continue 超时时间
}

func NewMultipleHostsReverseProxy(targets []*url.URL) *httputil.ReverseProxy {
	// 请求协调哲
	director := func(req *http.Request) {
		re, _ := regexp.Compile("^/dir(.*)")
		req.URL.Path = re.ReplaceAllString(req.URL.Path, "$1")
		// 随机负载均衡
		targetIndex := rand.Intn(len(targets))
		target := targets[targetIndex]
		targetQuery := target.RawQuery
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.Host = target.Host
		req.URL.Path = singleJoiningSlash(target.Path, req.URL.Path)
		if targetQuery == "" || req.URL.RawQuery == "" {
			req.URL.RawQuery = targetQuery + req.URL.RawQuery
		} else {
			req.URL.RawQuery = targetQuery + "&" + req.URL.RawQuery
		}
	}
	// 更改内容
	modifyFunc := func(resp *http.Response) error {
		if strings.Contains(resp.Header.Get("Connection"), "Upgrade") {
			return nil
		}
		var payload []byte
		var readErr error
		if strings.Contains(resp.Header.Get("Content-Encoding"), "gzip") {
			gr, err := gzip.NewReader(resp.Body)
			if err != nil {
				return err
			}
			payload, readErr = ioutil.ReadAll(gr)
			resp.Header.Del("Content-Encoding")
		} else {
			payload, readErr = ioutil.ReadAll(resp.Body)
		}
		if readErr != nil {
			return readErr
		}
		// 异常请求时设置StatusCode
		if resp.StatusCode != 200 {
			payload = []byte("StatusCode error:" + string(payload))
		}
		// NopCloser 将byte转换为Closer
		resp.Body = ioutil.NopCloser(bytes.NewBuffer(payload))
		// ContentLength 需要调整
		resp.ContentLength = int64(len(payload))
		// 需要在返回头加入 Content-Length
		resp.Header.Set("Content-Length", strconv.FormatInt(int64(len(payload)), 10))
		return nil
	}
	errFunc := func(w http.ResponseWriter, r *http.Request, err error) {
		http.Error(w, "ErrorHandler error:"+err.Error(), 500)
	}
	return &httputil.ReverseProxy{
		Director:       director,
		Transport:      transport,
		ModifyResponse: modifyFunc,
		ErrorHandler:   errFunc,
	}
}

func singleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}
