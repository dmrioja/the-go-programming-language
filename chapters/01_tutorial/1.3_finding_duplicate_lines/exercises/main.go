package main

import (
	"bufio"
	"fmt"
	"os"
)

// exercise_1_4
// Modify dup2 to print the names of all files in which
// each duplicated line occurs.
func main() {
	count := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countAndPrintFileName(os.Stdin, count)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Printf("dup2 -> error reading file %s: %v\n", arg, err)
				continue
			}
			countAndPrintFileName(f, count)
			f.Close()
		}
	}
}

func countAndPrintFileName(f *os.File, count map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		count[input.Text()]++
	}

	for _, n := range count {
		if n > 1 {
			fmt.Printf("file %s contains duplicated lines\n", f.Name())
			break
		}
	}
}
