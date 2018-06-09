package main

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRouting(t *testing.T) {
	tt := &[]struct {
		name       string
		method     string
		url        string
		statusCode int
		body       string
	}{
		{"index", http.MethodGet, "", http.StatusOK, "Hello World"},
		{"garbage", http.MethodGet, "/garbage", http.StatusOK, "Hello World"},
		{"method not get", http.MethodPost, "", http.StatusMethodNotAllowed, "only GET method allowed"},
	}

	s := httptest.NewServer(handler())
	defer s.Close()

	for _, tc := range *tt {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(tc.method, fmt.Sprintf("%s/%s", s.URL, tc.url), nil)
			if err != nil {
				t.Fatalf("could not %v req: %v", tc.method, err)
			}

			res, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Fatalf("could not get response: %v", err)
			}

			if res.StatusCode != tc.statusCode {
				t.Fatalf("expected status %v; got %v", tc.statusCode, res.Status)
			}

			b, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("could not read response: %v", err)
			}

			d := string(bytes.TrimSpace(b))
			if d != tc.body {
				t.Fatalf("expected %v; got %v", tc.body, d)
			}
		})
	}
}

func TestIndexContext(t *testing.T) {
	tt := &[]struct {
		name       string
		timeout    time.Duration
		statusCode int
	}{
		{"normal", 30, http.StatusOK},
		{"client timeout", 3, http.StatusInternalServerError},
	}

	for _, tc := range *tt {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, "/", nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()

			handler := http.HandlerFunc(indexHandler)

			ctx := req.Context()
			ctx, cancel := context.WithTimeout(ctx, tc.timeout*time.Second)
			defer cancel()

			req = req.WithContext(ctx)
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tc.statusCode {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tc.statusCode)
			}

		})
	}
}
