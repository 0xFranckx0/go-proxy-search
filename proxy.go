package main

import (
	r "github.com/0xFranckx0/go-proxy-search/pkg/rest"
	log "github.com/Sirupsen/logrus"
	"net/http"
	"os"
)

func main() {
	r.NewRouter()
	err := http.ListenAndServe(":"+os.Getenv("PORT"), r.NewRouter())
	if err != nil {
		log.Fatal(err)
	}
}
