package rest

import (
	//"encoding/json"
	"fmt"
	//log "github.com/Sirupsen/logrus"
	//"github.com/gorilla/mux"
	//"net/http"
)

/*
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.Methods("GET").Path("/version").Handler(http.HandlerFunc(versionHandler))
	return router
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	log.SetFormatter(&log.JSONFormatter{})

	type versionJson struct {
		Version string `json:"version"`
	}

	ver := versionJson{"1.0"}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(ver); err != nil {
		log.WithFields(log.Fields{
			"handler": "VersionHandler"}).Fatal(err)
	}
}
*/
func Toto() {
	fmt.Println("SALUT")
}
