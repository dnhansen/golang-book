package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			log.Fatalf("fetch: %v\n", err)
		}
		fmt.Println("Status code:", resp.Status)
		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			log.Fatalf("fetch: reading %s: %v\n", url, err)
		}
	}
}
