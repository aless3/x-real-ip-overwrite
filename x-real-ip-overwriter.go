// Package x-real-ip-overwrite
package aless3/x-real-ip-overwrite

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"text/template"
)

const (
	XRealIP        = "X-Real-IP"
	DefaultName    = "CF-Connecting-IP"
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
type XRealIPOverwriter struct {
	next         http.Handler
	headerName  string
	name         string
}

// Plugin:
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	if config.headerName == nil {
		return nil, fmt.Errorf("header name cannot be nil")
	}

	return &XRealIPOverwriter{
		next:       next,
    headerName: config.headerName,
		name:       name,
	}, nil
}

func (xrip *XRealIPOverwriter) ServeHTTP(rw http.ResponseWriter, req *http.Request) {


	ip := req.Header.Get(xrip.headerName)



	req.Header.Set("test", ip)

	xrip.next.ServeHTTP(rw, req)
}
