package rest

import (
	"encoding/json"
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/algolia/algoliasearch-client-go/algoliasearch"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var apiKey string

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

type handlerFunc func(http.ResponseWriter, *http.Request, *logrus.Entry) (int, string, []byte, string)

func (f handlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request, l *logrus.Entry) (int, string, []byte, string) {
	return f(w, r, l)
}

/*
	logger wrapper used for calling routeHandlers
*/
func logger(h handlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logFields := logrus.Fields{
			"method":    r.Method,
			"url":       r.URL.Path,
			"remotaddr": r.Header.Get("X-Real-IP"),
			"handler":   runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()}
		for key, val := range mux.Vars(r) {
			logFields[key] = val
		}
		log := logrus.WithFields(logFields)

		start := time.Now()
		code, contentType, data, msg := h.ServeHTTP(w, r, log)
		elapsed := time.Since(start)

		if contentType != "" {
			w.Header().Set("Content-Type", contentType)
		}
		w.WriteHeader(code)
		if data != nil {
			w.Write(data)
		}

		log = log.WithFields(logrus.Fields{"elapsedTime": elapsed.Seconds(), "code": code})
		if code >= 500 {
			log.Error(msg)
		} else {
			log.Info(msg)
		}
	})
}

/*
	StartRouter instantiates a mux Router and handle the routes defined below.
*/
func StartRouter() *mux.Router {
	router := mux.NewRouter()
	router.Methods("GET").Path("/version").Handler(logger(handlerFunc(versionHandler)))
	router.Methods("GET").Path("/1/search").Handler(logger(handlerFunc(searchHandler)))
	router.Methods("GET").Path("/1/usage/top_search").Handler(logger(handlerFunc(topSearchHandler)))

	return router
}

/*
	Implemented a version route just for convenience.
*/
func versionHandler(w http.ResponseWriter, r *http.Request, log *logrus.Entry) (int, string, []byte, string) {

	type versionJson struct {
		Version string `json:"version"`
	}

	ver := versionJson{"1.0"}

	respByte, err := json.Marshal(ver)
	if err != nil {
		return http.StatusInternalServerError, "", nil, "json.Marshal failed: " + err.Error()
	}

	return http.StatusOK, "application/json", respByte, ""
}

/*
	SearchHanlder is aimed to perfom a search on the best_buy index:

	usage:
		curl -XGET -v  http://localhost:${PORT}/1/search?query=price
		curl -XGET -v https://calm-lowlands-40938.herokuapp.com/1/search?query=price

	You need to provide your Algolia Public API KEY as variable
*/
func searchHandler(w http.ResponseWriter, r *http.Request, log *logrus.Entry) (int, string, []byte, string) {
	params := r.URL.Query()
	query := params["query"][0]
	appId := "AVGXBPGE8G"
	publicApiKey := "900876a37773a1ea2f6b56fda68d0518"
	logrus.Info("starting search")

	//use the algolia client to perform search on our index
	client := algoliasearch.NewClient(appId, publicApiKey)
	index := client.InitIndex("best_buy")

	res, err := index.Search(query, nil)
	if err != nil {
		return http.StatusInternalServerError, "", nil, "index.Search failed: " + err.Error()
	}
	if res.Hits == nil {
		return http.StatusInternalServerError, "", nil, "Missing hits"

	}

	//Extract hits attribute to just return it
	type result struct {
		Hits interface{} `json:"hits"`
	}
	rslt := result{Hits: res.Hits}
	respByte, err := json.Marshal(rslt)
	if err != nil {
		return http.StatusInternalServerError, "", nil, "json.Marshal failed: " + err.Error()
	}

	return http.StatusOK, "application/json", respByte, ""
}

/*
	topSearchHanlder is aimed to perfom a popular search a described by :
	https://www.algolia.com/doc/rest-api/analytics/#get-popular-searches

	usage:
		curl -XGET -v  http://localhost:${PORT}/1/usage/top_search?size=10
		curl -XGET -v https://calm-lowlands-40938.herokuapp.com/1/usage/top_search?size=10

	You need to provide your Algolia admin API KEY. Store it in a file /app/keyfile
*/
func topSearchHandler(w http.ResponseWriter, r *http.Request, log *logrus.Entry) (int, string, []byte, string) {

	params := r.URL.Query()
	size, err := strconv.Atoi(params["size"][0])
	if err != nil {
		return http.StatusInternalServerError, "", nil, "strconv.Atoi failed: " + err.Error()
	}
	appId := "AVGXBPGE8G"
	if len(apiKey) == 0 {
		fmt.Println("SET KEY")
		buf, err := ioutil.ReadFile("./keyfile")
		if err != nil {
			return http.StatusInternalServerError, "", nil, "ioutil.ReadFile failed: " + err.Error()
		}
		apiKey = string(buf)
		apiKey = strings.Trim(apiKey, "\n")
	}

	method := "GET"
	path := `/1/searches/best_buy/popular`
	host := "analytics.algolia.com"
	url := "https://" + host + path

	headers := map[string]string{
		"Connection":               "keep-alive",
		"User-Agent":               "REST Proxy Go ",
		"X-Algolia-API-Key":        apiKey,
		"X-Algolia-Application-Id": appId,
	}

	// build our http request
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return http.StatusInternalServerError, "", nil, "http.NewRequest failed: " + err.Error()
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	resp, err := client.Do(req)
	if err != nil {
		return http.StatusInternalServerError, "", nil, "client.Do failed: " + err.Error()
	}
	defer resp.Body.Close()

	// We need to just keep the number of <size> elements
	type top struct {
		TopSearches []interface{} `json:"topsearches"`
	}

	var result top
	var body map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		return http.StatusInternalServerError, "", nil, "client.Do failed: " + err.Error()
	}

	// keep track if we have a topsearch key in the returned JSON
	h := false

	// iterate until we match the topSearch field
	for key, value := range body {
		if key == "topSearches" {
			h = true

			// iterate to just keep the <size> top elements
			for _, x := range value.([]interface{}) {
				if size == 0 {
					break
				}
				size--
				result.TopSearches = append(result.TopSearches, x)
			}
		}
	}

	if !h {
		return http.StatusInternalServerError, "", nil, "topSearches Not Found: " + err.Error()
	}

	// encode the result as json
	data, err := json.Marshal(result)
	if err != nil {
		return http.StatusInternalServerError, "", nil, "ioutil.ReadAll failed: " + err.Error()
	}

	return http.StatusOK, "application/json", data, ""
}
