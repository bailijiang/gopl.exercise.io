package main

import (
	"bufio"
	"fmt"
	"os"
)

// Modify dup2 to print the names of all files in which each duplicated line occurs.
func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		// Put empty string because there is no file.
		countLines("", os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(arg, f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(n string, f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[n+": "+input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
