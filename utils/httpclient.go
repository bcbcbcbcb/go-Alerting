package utils

import (
	"net"
	"net/http"
	"net/url"
	"time"
)

var (
	_httpClient *http.Client
)

func init() {
	config_init()
	// netproxy := "192.168.160.18:3129"
	// netproxy := "nil"
	netproxy := Config.GetString("http.proxy")
	proxy := func(_ *http.Request) (*url.URL, error) {
		if netproxy != "nil" {
			return url.Parse("http://" + netproxy)
		} else {
			return nil, nil
		}
	}
	httpTransport := &http.Transport{
		Proxy: proxy,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	_httpClient = &http.Client{
		Transport: httpTransport,
		Timeout:   100 * time.Second,
	}
}

func GetHttpCilent() *http.Client {
	return _httpClient
}
