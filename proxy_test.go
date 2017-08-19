package main_test

import (
	"encoding/json"
	"github.com/0xFranckx0/go-proxy-search/pkg/rest"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestVersion(t *testing.T) {
	req, _ := http.NewRequest("GET", "/version", nil)
	resp := sendReq(req)

	chkRespCode(t, http.StatusOK, resp.Code)

	if body := resp.Body.String(); body != "" {
		type versionJson struct {
			Version string `json:"version"`
		}

		var j versionJson

		err := json.Unmarshal([]byte(body), &j)
		if err != nil {
			t.Errorf("%s", err)
		}

		if j.Version != "1.0" {
			t.Errorf("BAD VERSION %s", j.Version)
		}
	}
}

func sendReq(req *http.Request) *httptest.ResponseRecorder {
	rec := httptest.NewRecorder()
	router := rest.StartRouter()
	router.ServeHTTP(rec, req)

	return rec
}

func chkRespCode(t *testing.T, ok, code int) {
	if ok != code {
		t.Errorf("Want %d. Got %d\n", ok, code)
	}
}
