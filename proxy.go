package main

import (
	"github.com/0xFranckx0/go-proxy-search/pkg/rest"
	log "github.com/Sirupsen/logrus"
	"net/http"
	"os"
)

/*

This is the main funtion that starts the API server.
The port on which the API server listens is set by the
PORT environment variable.

*/
func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.Info("starting proxy")

	router := rest.StartRouter()
	port := ":" + os.Getenv("PORT")

	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Fatal(err)
	}
}
