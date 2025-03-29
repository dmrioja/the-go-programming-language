package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	exercise_1_10()
	exercise_1_11()
}

// exercise_1_10
// Find a web site that produces a large amount of data.
// Investigate caching by running fetchall twice in succession to see
// wether the reported time changes much. Do you get the same content
// each time? Modify fecthall to print its output to a file so it can
// be examined.
func exercise_1_10() {
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		for i := 0; i < 4; i++ {
			go fetch(url, ch, i)
		}
	}

	for range 4 {
		fmt.Printf(<-ch) // receive from channel ch
		/*
			0.95s    137497         https://pkg.go.dev/fmt
			0.27s    137497         https://pkg.go.dev/fmt
			0.20s    137497         https://pkg.go.dev/fmt
			0.20s    137497         https://pkg.go.dev/fmt
		*/
	}
}

// exercise_1_11
// Try fetchall with longer argument lists, such as samples from the
// top million web sites available at alexa.com. How does the program
// behave if a web site just doesn't respond?
func exercise_1_11() {
	// Basically same than exercise 1.10
}

func fetch(url string, ch chan<- string, count int) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	fileName := fmt.Sprintf("exercise_1_10_output_%d.txt", count)
	file, err := os.Create(fmt.Sprintf("%s/%s", "chapters/01_tutorial/1.6_fetching_urls_concurrently/exercises", fileName))
	if err != nil {
		ch <- fmt.Sprintf("creating file %s: %v", fileName, err)
		return
	}
	defer file.Close()

	nbytes, err := io.Copy(file, resp.Body)
	resp.Body.Close() // don`t leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	ch <- fmt.Sprintf("%.2fs\t%7d\t\t%s\n", time.Since(start).Seconds(), nbytes, url)
}
