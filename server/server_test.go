package server

import (
	"context"
	"io"
	"kvdb/db"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
)

func TestServer(t *testing.T) {
	tests := map[string]func(t *testing.T){
		"TestPutAndGet": func(t *testing.T) {
			store := db.New()
			srv := NewServer(store)
			server := httptest.NewServer(srv.Route())
			defer server.Close()

			client := &http.Client{}

			// PUT request
			req, _ := http.NewRequest(http.MethodPut, server.URL+"/k", strings.NewReader("v"))
			resp, err := client.Do(req)
			if err != nil || resp.StatusCode != http.StatusOK {
				t.Fatalf("PUT request failed: %v", err)
			}

			// GET request
			req, _ = http.NewRequest(http.MethodGet, server.URL+"/k", nil)
			resp, err = client.Do(req)
			if err != nil || resp.StatusCode != http.StatusOK {
				t.Fatalf("GET request failed: %v", err)
			}
			body, _ := io.ReadAll(resp.Body)
			if string(body) != "v" {
				t.Errorf("Expected 'v', got '%s'", string(body))
			}
		},

		"TestDelete": func(t *testing.T) {
			store := db.New()
			ctx := context.Background()
			store.Set(ctx, "k", "v")
			srv := NewServer(store)
			server := httptest.NewServer(srv.Route())
			defer server.Close()

			client := &http.Client{}

			req, _ := http.NewRequest(http.MethodDelete, server.URL+"/k", nil)
			resp, err := client.Do(req)
			if err != nil || resp.StatusCode != http.StatusOK {
				t.Fatalf("DELETE request failed: %v", err)
			}

			// confirm deletion
			req, _ = http.NewRequest(http.MethodGet, server.URL+"/k", nil)
			resp, err = client.Do(req)
			if err != nil {
				t.Fatalf("GET request failed: %v", err)
			}
			if resp.StatusCode != http.StatusNotFound {
				t.Errorf("Expected 404 Not Found, got %d", resp.StatusCode)
			}
		},

		"TestListKeys": func(t *testing.T) {
			store := db.New()
			ctx := context.Background()
			store.Set(ctx, "k1", "v1")
			store.Set(ctx, "k2", "v2")
			srv := NewServer(store)
			server := httptest.NewServer(srv.Route())
			defer server.Close()

			client := &http.Client{}

			req, _ := http.NewRequest(http.MethodGet, server.URL+"/", nil)
			resp, err := client.Do(req)
			if err != nil || resp.StatusCode != http.StatusOK {
				t.Fatalf("GET request failed: %v", err)
			}
			body, _ := io.ReadAll(resp.Body)

			expected1 := `["k1","k2"]`
			expected2 := `["k2","k1"]`

			bodyStr := strings.TrimSpace(string(body))
			if bodyStr != expected1 && bodyStr != expected2 {
				t.Errorf("Expected %s or %s, got %s", expected1, expected2, bodyStr)
			}
		},

		"TestParallelPuts": func(t *testing.T) {
			store := db.New()
			srv := NewServer(store)
			server := httptest.NewServer(srv.Route())
			defer server.Close()

			client := &http.Client{}
			var wg sync.WaitGroup
			values := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}

			for _, val := range values {
				wg.Add(1)
				go func(value string) {
					defer wg.Done()
					req, _ := http.NewRequest(http.MethodPut, server.URL+"/key", strings.NewReader(value))
					resp, err := client.Do(req)
					if err != nil || resp.StatusCode != http.StatusOK {
						t.Errorf("PUT request failed: %v", err)
					}
				}(val)
			}
			wg.Wait()

			req, _ := http.NewRequest(http.MethodGet, server.URL+"/key", nil)
			resp, err := client.Do(req)
			if err != nil || resp.StatusCode != http.StatusOK {
				t.Fatalf("GET request failed: %v", err)
			}
			body, _ := io.ReadAll(resp.Body)
			value := string(body)
			found := false
			for _, v := range values {
				if v == value {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("Final value '%s' not in expected values", value)
			}
		},
	}

	for name, testFunc := range tests {
		t.Run(name, func(t *testing.T) {
			testFunc(t)
		})
	}
}
