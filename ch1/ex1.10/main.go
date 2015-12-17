// Find a web site that produces a large amount of data. Investigate caching by
// running fetchall twice in scuccession to see whether the reported time
// changes much. Do you get the same content each time? Modify fetch all to
// print its output to a file so it can be examined.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	// Use domain as a file name becuase a file name with protocol
	// causes an error when it is created.
	domain := strings.Split(url, "://")[1]
	file, err := os.Open(domain)
	// Fetch URL if the file doesn't exist.
	if err != nil {
		fmt.Println(err)
		fmt.Println("Create a file.", domain)

		file, err = os.Create(domain)
		if err != nil {
			fmt.Println(err)
			ch <- fmt.Sprint(err)
			return
		}
		resp, err := http.Get(url)
		if err != nil {
			ch <- fmt.Sprint(err)
			return
		}
		// Ignore the byte size. Maybe bad idea.
		_, err = io.Copy(file, resp.Body)
		resp.Body.Close()
		if err != nil {
			ch <- fmt.Sprintf("while reading %s: %v", url, err)
			return
		}
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %s", secs, url)
}
