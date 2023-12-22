package x_real_ip_overwrite

import (
	"context"
	"fmt"
	"net/http"
)

const (
	XRealIP        = "X-Real-IP"
	DefaultName    = "CF-Connecting-IP"
	OverwrittenName    = "X-Real-IP-Overwritten"
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
	if config.headerName == "" {
		return nil, fmt.Errorf("header name cannot be empty")
	}

	return &XRealIPOverwriter{
		next:       next,
    headerName: config.headerName,
		name:       name,
	}, nil
}

func (xrip *XRealIPOverwriter) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	ip := req.Header.Get(xrip.headerName)
  if ip != "" {
    realWrong := req.Header.Get(XRealIP)
    req.Header.Set(OverwrittenName, realWrong)
    req.Header.Set(XRealIP, ip)
  }

	xrip.next.ServeHTTP(rw, req)
}
