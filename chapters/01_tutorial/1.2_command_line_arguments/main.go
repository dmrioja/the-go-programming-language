package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	echo1()
	echo2()
	echo3()
}

// echo1 is the first version of Unix echo command
func echo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

// echo2 is the second version of Unix echo command
func echo2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

// echo3 is the third version of Unix echo command
func echo3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
