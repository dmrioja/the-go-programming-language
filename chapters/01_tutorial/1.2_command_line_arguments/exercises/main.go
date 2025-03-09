package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	exercise_1_1()
	exercise_1_2()
	exercise_1_3()
}

// exercise_1_1
// Modify the echo program to also print os.Args[0], the name
// of the command that invoked it.
func exercise_1_1() {
	fmt.Println(strings.Join(os.Args, " "))
}

// exercise_1_2
// Modify the echo program to print the index and value of each
// of its arguments, one per line.
func exercise_1_2() {
	for index, value := range os.Args[1:] {
		fmt.Println(index, "-", value)
	}
}

// exercise_1_3
// Experiment to measure the difference in running time between our
// potentially inefficient versions and the one that uses strings.Join.
func exercise_1_3() {
	measureTime(echo1)
	measureTime(echo2)
	measureTime(echo3)
}

func measureTime(f func()) {
	start := time.Now()
	f()
	fmt.Println(time.Since(start))
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
