package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)
	os.Args[1] // give me everything after that first "temp file path"
}
