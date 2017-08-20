.Phony: install build deps clean 

install: build
	go install

build: deps
	go build -o proxy proxy.go 
deps: clean
	go get golang.org/x/sys/unix
	go get github.com/algolia/algoliasearch-client-go/algoliasearch
	go get github.com/Sirupsen/logrus
	go get golang.org/x/crypto/ssh/terminal
	go get github.com/gorilla/mux
	govendor init
	govendor fetch github.com/0xFranckx0/go-proxy-search/pkg/rest
	govendor fetch golang.org/x/crypto/ssh/terminal
	godep save 
clean:
	-rm -rf vendor
	-rm -rf Godep 
