package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("Water freezing temperature:\t%gºF = %gºC\n", freezingF, fToC(freezingF))
	fmt.Printf("Water boiling temperature:\t%gºF = %gºC\n", boilingF, fToC(boilingF))
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
