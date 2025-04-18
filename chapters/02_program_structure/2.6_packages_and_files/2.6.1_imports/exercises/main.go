package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/dmrioja/the-go-programming-language/chapters/02_program_structure/2.6_packages_and_files/lengthconv"
	"github.com/dmrioja/the-go-programming-language/chapters/02_program_structure/2.6_packages_and_files/tempconv"
	"github.com/dmrioja/the-go-programming-language/chapters/02_program_structure/2.6_packages_and_files/weightconv"
)

func main() {

	args := os.Args[1:]

	if len(args) == 0 {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			printMeasuresValues(input.Text())
		}
	} else {
		for _, arg := range args {
			printMeasuresValues(arg)
		}
	}
}

func printMeasuresValues(input string) {
	value := extractFloat64Value(input)
	printTemperatureValues(value)
	printLengthValues(value)
	printWeightValues(value)
}

func extractFloat64Value(s string) float64 {
	value, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	return value
}

func printTemperatureValues(value float64) {
	c := tempconv.Celsius(value)
	f := tempconv.Fahrenheit(value)
	fmt.Println("TEMPERATURE")
	fmt.Printf("%v equals %v\n", c, tempconv.CToF(c))
	fmt.Printf("%v equals %v\n\n", f, tempconv.FToC(f))
}

func printLengthValues(value float64) {
	m := lengthconv.Meter(value)
	ft := lengthconv.Feet(value)
	fmt.Println("LENGTH")
	fmt.Printf("%v equals %v\n", m, lengthconv.MToFt(m))
	fmt.Printf("%v equals %v\n\n", ft, lengthconv.FtToM(ft))
}

func printWeightValues(value float64) {
	kg := weightconv.Kilogram(value)
	lbs := weightconv.Pound(value)
	fmt.Println("WEIGHT")
	fmt.Printf("%v equals %v\n", kg, weightconv.KgToLbs(kg))
	fmt.Printf("%v equals %v\n\n", lbs, weightconv.LbsToKg(lbs))
}
