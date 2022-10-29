package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type set map[string]bool

func main() {
	counts := make(map[string]int)
	files := make(map[string]set)

	filenames := os.Args[1:]
	if len(filenames) == 0 {
		countLines(os.Stdin, counts, files)
	} else {
		for _, arg := range filenames {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, files)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, files[line].string())
		}
	}
}

func countLines(f *os.File, counts map[string]int, files map[string]set) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if _, ok := files[input.Text()]; !ok {
			files[input.Text()] = make(set)
		}
		files[input.Text()][f.Name()] = true

	}
	// NOTE: ignoring potential errors from input.Err()
}

func (s set) string() string {
	i := 0
	elems := make([]string, len(s))
	for k := range s {
		elems[i] = k
		i++
	}

	return strings.Join(elems, ", ")
}
