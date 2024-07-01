package google_translation

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"
)

var transport = &http.Transport{
	DisableKeepAlives:  true,
	DisableCompression: true,
	MaxConnsPerHost:    3,
	DialContext: (&net.Dialer{
		Timeout:   10 * time.Second,
		KeepAlive: 60 * time.Second,
	}).DialContext,
	ExpectContinueTimeout: 5 * time.Second,
	ResponseHeaderTimeout: 5 * time.Second,
	// We use ABSURDLY large keys, and should probably not.
	TLSHandshakeTimeout: 5 * time.Second,
	TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
}

var client = &http.Client{
	Timeout:   30 * time.Second,
	Transport: transport,
}
