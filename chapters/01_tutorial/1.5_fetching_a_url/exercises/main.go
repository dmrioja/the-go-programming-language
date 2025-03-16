package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	exercise_1_7()
	exercise_1_8()
	exercise_1_9()
}

// exercise_1_7
// The function call io.Copy(dst, src) reads from src and writes to dst.
// Use it instead of io.ReadAll to copy the response body to os.Stdout without
// requiring a buffer large enough to hold the entire stream. Be sure to check
// the error result of io.Copy.
func exercise_1_7() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("fetch err: %v\n", err)
			os.Exit(1)
		}

		b, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Printf("fetch err reading %s: %v\n", url, err)
			os.Exit(1)
		}

		fmt.Printf("URL %s fetched %d bytes\n", url, b)
	}
}

// exercise_1_8
// Modify fetch to add the prefix http:// to each argument URL if it is missing.
// You might want to use strings.HasPrefix
func exercise_1_8() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("fetch err: %v\n", err)
			os.Exit(1)
		}

		b, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Printf("fetch err reading %s: %v\n", url, err)
			os.Exit(1)
		}

		fmt.Printf("%s\n", b)
	}
}

// exercise_1_9
// Modify fetch to also print the HTTP status code, found in resp.Status.
func exercise_1_9() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("fetch err: %v\n", err)
			os.Exit(1)
		}

		b, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Printf("fetch %s responded with [%d] StatusCode but got err reading body: %v\n", url, resp.StatusCode, err)
			os.Exit(1)
		}

		fmt.Printf("fetch %s responded with [%d] StatusCode and body:\n%s\n", url, resp.StatusCode, b)
	}
}
