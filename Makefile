.Phony: build deps


build: deps
	go build -o proxy cmd/proxy.go 
deps:
	go get ./... 
