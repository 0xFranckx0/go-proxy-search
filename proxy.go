package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Getenv("MSG"))
	fmt.Println("Hello World!")

}
