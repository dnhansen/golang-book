// NOTE: To run, do
//    go run *.go
// in terminal.

package main

import (
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	cycles := r.FormValue("cycles")
	c, err := strconv.Atoi(cycles)
	if err != nil {
		c = 5 // Default value
	}

	lissajous(w, float64(c))
}
