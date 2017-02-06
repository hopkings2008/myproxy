package main

import (
	"encoding/base64"
	"log"
	"net/http"
	"strings"

	"github.com/elazarl/goproxy"
)

func main() {
	proxy := goproxy.NewProxyHttpServer()
	proxy.OnRequest().DoFunc(
		func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
			user, pwd, ok := proxyBasicAuth(r)
			if !ok {
				return r, goproxy.NewResponse(r,
					goproxy.ContentTypeText, http.StatusForbidden, "forbidden")
			}
			if user != "shiqiuser" || pwd != "test123" {
				return r, goproxy.NewResponse(r,
					goproxy.ContentTypeText, http.StatusForbidden, "forbidden")
			}
			return r, nil
		})
	proxy.Verbose = true
	log.Fatal(http.ListenAndServe(":9999", proxy))
}

func proxyBasicAuth(r *http.Request) (string, string, bool) {
	user, pwd, ok := r.BasicAuth()
	if ok {
		return user, pwd, ok
	}
	auth := r.Header.Get("Proxy-Authorization")
	if auth == "" {
		return "", "", false
	}
	return parseBasicAuth(auth)
}

func parseBasicAuth(auth string) (username, password string, ok bool) {
	const prefix = "Basic "
	if !strings.HasPrefix(auth, prefix) {
		return
	}
	c, err := base64.StdEncoding.DecodeString(auth[len(prefix):])
	if err != nil {
		return
	}
	cs := string(c)
	s := strings.IndexByte(cs, ':')
	if s < 0 {
		return
	}
	return cs[:s], cs[s+1:], true
}
