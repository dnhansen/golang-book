package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args {
		fmt.Printf("Argument #%d: %s\n", i, arg)
	}
}