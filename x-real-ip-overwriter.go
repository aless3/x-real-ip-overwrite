package xrip

import (
	"context"
	"fmt"
	"net/http"
)

const (
	XRealIP     = "X-Real-IP"
	DefaultName = "CF-Connecting-IP"
	XFF         = "X-Forwarded-For"
	Overwritten = "X-Real-IP-Overwritten"
)

// Config the plugin configuration.
type Config struct {
	headerName string `json:"header-name" toml:"header-name" yaml:"header-name"`
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		headerName: DefaultName,
	}
}

// Demo a Demo plugin.
type XRIPOverwrite struct {
	next       http.Handler
	headerName string
	name       string
}

// Plugin:
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	if config.headerName == "" {
		return nil, fmt.Errorf("header name cannot be empty")
	}

	return &XRIPOverwrite{
		next:       next,
		headerName: config.headerName,
		name:       name,
	}, nil
}

func (xrip *XRIPOverwrite) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	ip := req.Header.Get(xrip.headerName)
	if ip != "" {
		XRIPWrong := req.Header.Get(XRealIP)
		req.Header.Set(Overwritten, XRIPWrong)

		req.Header.Set(XRealIP, ip)
		req.Header.Set(XFF, ip)
	}

	xrip.next.ServeHTTP(rw, req)
}
