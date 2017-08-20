.Phony: install build deps local-deps


install: build
	go install

build: deps
	go build -o proxy proxy.go 
deps:
	go get golang.org/x/sys/unix
	go get github.com/algolia/algoliasearch-client-go/algoliasearch
	go get ./... 

local-deps:
	govendor fetch github.com/0xFranckx0/go-proxy-search/pkg/rest
	godep save ./...
