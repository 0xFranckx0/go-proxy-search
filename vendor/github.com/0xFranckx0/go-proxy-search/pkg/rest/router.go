package rest

import (
	"encoding/json"
	"github.com/Sirupsen/logrus"
	"github.com/algolia/algoliasearch-client-go/algoliasearch"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"os"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func StartRouter() *mux.Router {
	router := mux.NewRouter()
	router.Methods("GET").Path("/1/search").Handler(http.HandlerFunc(searchHandler))
	router.Methods("GET").Path("/1/usage/top_search").Handler(http.HandlerFunc(topSearchHandler))

	return router
}

/*
	SearchHanlder is aimed to perfomed a search on the best_index:

	usage:
		curl -XGET -v  http://localhost:8181/1/search?query=price
		curl -XGET -v https://calm-lowlands-40938.herokuapp.com/1/search?query=price

	You need to provide your Algolia Public API KEY as variable
*/
func searchHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	query := params["query"][0]
	appId := "AVGXBPGE8G"
	publicApiKey := "900876a37773a1ea2f6b56fda68d0518"
	logrus.Info("starting search")

	client := algoliasearch.NewClient(appId, publicApiKey)
	index := client.InitIndex("best_buy")

	res, err := index.Search(query, nil)
	if err != nil {
		logrus.Error(err)
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(res); err != nil {
		logrus.WithFields(logrus.Fields{
			"handler": "VersionHandler"}).Fatal(err)
	}
}

/*
	topSearchHanlder is aimed to perfomed a popular search a described by :
	https://www.algolia.com/doc/rest-api/analytics/#get-popular-searches

	usage:
		curl -XGET -v  http://localhost:8181/1/usage/top_search?size=10
		curl -XGET -v https://calm-lowlands-40938.herokuapp.com/1/usage/top_search?size=10

	You need to provide your Algolia admin API KEY as variable environment.
*/
func topSearchHandler(w http.ResponseWriter, r *http.Request) {
	logrus.WithFields(logrus.Fields{
		"handler": "usageHandler"}).Info("Starting top search")

	params := r.URL.Query()
	size := params["size"][0]
	appId := "AVGXBPGE8G"
	adminApiKey := os.Getenv("ADMIN_API_KEY")
	method := "GET"
	path := `/1/searches/best_buy/popular`
	host := "analytics.algolia.com"
	url := "https://" + host + path

	headers := map[string]string{
		"Connection":               "keep-alive",
		"User-Agent":               "REST Proxy Go ",
		"X-Algolia-API-Key":        adminApiKey,
		"X-Algolia-Application-Id": appId,
	}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"handler": "topSearchHandler"}).Error(err)
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	q := req.URL.Query()
	q.Add("size", size)

	req.URL.RawQuery = q.Encode()

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"handler": "topSearchHandler"}).Error(err)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
