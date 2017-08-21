package main_test

import (
	"encoding/json"
	"github.com/0xFranckx0/go-proxy-search/pkg/rest"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSearch(t *testing.T) {
	fetchApi(t, "GET", "/1/search", "query", "price", "hits")
}

func TestTopSearch(t *testing.T) {
	fetchApi(t, "GET", "/1/usage/top_search", "size", "10", "topsearches")
}

func fetchApi(t *testing.T, method string, path string, key string, value string, pattern string) {
	req, _ := http.NewRequest(method, path, nil)

	q := req.URL.Query()
	q.Add(key, value)

	req.URL.RawQuery = q.Encode()

	resp := httptest.NewRecorder()
	router := rest.StartRouter()
	router.ServeHTTP(resp, req)

	if http.StatusOK != resp.Code {
		t.Errorf("NOT OK %d\n", resp.Code)
	}

	if body := resp.Body.String(); body != "" {
		c := make(map[string]interface{})

		err := json.Unmarshal([]byte(body), &c)
		if err != nil {
			t.Errorf("%s", err)
		}
		h := false

		for s, _ := range c {
			if s == pattern {
				h = true
			}
		}
		if !h {
			t.Errorf("MISSING %s", pattern)

		}
	}

}
