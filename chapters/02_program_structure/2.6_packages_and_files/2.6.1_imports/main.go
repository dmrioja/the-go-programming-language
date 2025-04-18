package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/dmrioja/the-go-programming-language/chapters/02_program_structure/2.6_packages_and_files/tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		value, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cfk: %v\n", err)
			os.Exit(1)
		}

		f := tempconv.Fahrenheit(value)
		c := tempconv.Celsius(value)
		k := tempconv.Kelvin(value)

		fmt.Printf("Input value: %g\n", value)
		fmt.Printf("FAHRENHEIT: %s equal %s or %s\n", f, tempconv.FToC(f), tempconv.FToK(f))
		fmt.Printf("CELSIUS: %s equal %s or %s\n", c, tempconv.CToF(c), tempconv.CToK(c))
		fmt.Printf("KELVIN: %s equal %s or %s\n\n", k, tempconv.KToC(k), tempconv.KToF(k))
	}
}
