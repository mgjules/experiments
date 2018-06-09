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
		url        string
		statusCode int
		body       string
	}{
		{"index", "", http.StatusOK, "Hello World"},
		{"garbage", "/garbage", http.StatusOK, "Hello World"},
	}

	s := httptest.NewServer(handler())
	defer s.Close()

	for _, tc := range *tt {
		t.Run(tc.name, func(t *testing.T) {
			res, err := http.Get(fmt.Sprintf("%s/%s", s.URL, tc.url))
			if err != nil {
				t.Fatalf("could not GET req: %v", err)
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
		method     string
		timeout    time.Duration
		statusCode int
	}{
		{"normal", http.MethodGet, 30, http.StatusOK},
		{"client timeout", http.MethodGet, 3, http.StatusInternalServerError},
		{"method not get", http.MethodPost, 30, http.StatusMethodNotAllowed},
	}

	for _, tc := range *tt {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(tc.method, "/", nil)
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
