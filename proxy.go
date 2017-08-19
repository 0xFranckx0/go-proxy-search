package main

import (
	"fmt"
	"github.com/0xFranckx0/go-proxy-search/pkg/rest"
	"os"
)

func main() {
	fmt.Println(os.Getenv("MSG"))
	fmt.Println("Hello World!")
	rest.Test()
}
