package x_real_ip_overwrite

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	// "github.com/aless3/x-real-ip-overwrite"
)

// no changes
func TestDemoNoCF(t *testing.T) {
	cfg := XRealIPOverwriter.CreateConfig()

	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := x-real-ip-overwrite.New(ctx, next, cfg, "x-real-ip-overwrite")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Headers["X-Real-IP"] = "127.0.0.1"
	handler.ServeHTTP(recorder, req)

	assertHeader(t, req, "X-Real-IP", "127.0.0.1")
}


func assertHeader(t *testing.T, req *http.Request, key, expected string) {
	t.Helper()

	if req.Header.Get(key) != expected {
		t.Errorf("invalid header value: %s", req.Header.Get(key))
	}
}