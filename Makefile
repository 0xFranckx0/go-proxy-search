.Phony: install build deps heroku-goapp test

default: install

install: deps 
	go install

build: 
	go build -o proxy proxy.go 

deps: 
	go get golang.org/x/sys/unix
	go get github.com/algolia/algoliasearch-client-go/algoliasearch
	go get github.com/Sirupsen/logrus
	go get golang.org/x/crypto/ssh/terminal
	go get github.com/gorilla/mux
	go get github.com/0xFranckx0/go-proxy-search/pkg/rest
	cp pkg/rest/router.go vendor/github.com/0xFranckx0/go-proxy-search/pkg/rest/router.go

heroku-goapp:	
	govendor init
	govendor fetch github.com/0xFranckx0/go-proxy-search/pkg/rest
	govendor fetch golang.org/x/crypto/ssh/terminal
	godep save 

test:
	go test -run TestSearch 
