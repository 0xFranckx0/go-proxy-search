package main

import (
	"fmt"
	r "github.com/0xFranckx0/go-proxy-search/pkg/rest"
	log "github.com/Sirupsen/logrus"
	"net/http"
	"os"
)

func main() {
	fmt.Println(os.Getenv("MSG"))
	fmt.Println("Hello World!")
	r.Toto()
	r.NewRouter()
	err := http.ListenAndServe(":80", r.NewRouter())
	if err != nil {
		log.Fatal(err)
	}
}
