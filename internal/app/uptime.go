package app

import (
	"crypto/tls"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"go.uber.org/zap"
	"net"
	"net/http"
	"time"
)

var client = &http.Client{
	Timeout: 30 * time.Second,
	Transport: &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   60 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	},
}

func UptimeHandler() gin.HandlerFunc {
	start := time.Now()
	return func(c *gin.Context) {
		req, err := http.NewRequestWithContext(c, "GET", fmt.Sprintf("https://img.shields.io/badge/%s-brightgreen", time.Since(start).Truncate(time.Second).String()), nil)
		if err != nil {
			c.String(400, err.Error())
			log.Error("", zap.Error(err))
			return
		}
		req.Header["Accept"] = []string{"text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"}
		req.Header["Accept-Language"] = []string{"en-US,en;q=0.9,ru;q=0.8"}
		req.Header["User-Agent"] = []string{"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.41 Safari/537.36"}

		resp, err := client.Do(req)
		if err != nil {
			c.String(400, err.Error())
			log.Error("", zap.Error(err))
			return
		}
		defer resp.Body.Close()
		c.DataFromReader(200, resp.ContentLength, "image/svg+xml;charset=utf-8", resp.Body, nil)
	}
}
