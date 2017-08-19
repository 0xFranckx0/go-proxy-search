package rest

import (
	"encoding/json"
	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"net/http"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func StartRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.Methods("GET").Path("/version").Handler(http.HandlerFunc(versionHandler))

	return router
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	logrus.WithFields(logrus.Fields{
		"handler": "VersionHandler"}).Info("VERSION")
	type versionJson struct {
		Version string `json:"version"`
	}

	ver := versionJson{"1.0"}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(ver); err != nil {
		logrus.WithFields(logrus.Fields{
			"handler": "VersionHandler"}).Fatal(err)
	}
}
