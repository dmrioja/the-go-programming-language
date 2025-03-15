package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	dup1()
	dup2()
	dup3()
}

func dup1() {
	count := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		count[input.Text()]++
	}

	printLines(count)
}

func dup2() {
	count := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countAndPrintLines(os.Stdin, count)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Printf("dup2 -> error reading file %s: %v\n", arg, err)
				continue
			}
			countAndPrintLines(f, count)
			f.Close()
		}
	}
}

func countAndPrintLines(f *os.File, count map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		count[input.Text()]++
	}

	printLines(count)
}

func dup3() {
	count := make(map[string]int)
	for _, arg := range os.Args[1:] {
		data, err := os.ReadFile(arg)
		if err != nil {
			fmt.Printf("dup3 -> error reading file %s: %v\n", arg, err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			count[line]++
		}
	}

	printLines(count)
}

func printLines(count map[string]int) {
	for line, n := range count {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
