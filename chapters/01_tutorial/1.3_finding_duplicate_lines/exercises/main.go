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
	counter := make(map[string][]string)
	files := os.Args[1:]

	// I will only consider reading files an not Stdin
	for _, arg := range files {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Printf("dup2 -> error reading file %s: %v\n", arg, err)
			continue
		}
		count := countDupLines(f)
		f.Close()

		for line, n := range count {
			if n > 1 {
				counter[line] = append(counter[line], arg)
			}
		}
	}

	for line, files := range counter {
		fmt.Printf("line \"%s\" appears duplicated in files %v\n", line, files)
	}
}

func countDupLines(f *os.File) map[string]int {
	count := make(map[string]int)
	input := bufio.NewScanner(f)
	for input.Scan() {
		count[input.Text()]++
	}

	return count
}
