.Phony: build deps


build: deps
	go build -o proxy proxy.go 
deps:
	go get ./... 
