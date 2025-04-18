package main

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100

	AbsoluteZeroF Fahrenheit = -459.67
	FreezingF     Fahrenheit = 32
	BoilingF      Fahrenheit = 212
)

func main() {

	fmt.Println("Converting ºCelsius to ºFahrenheit")
	fmt.Printf("Absolute zero: %v or %v\n", AbsoluteZeroC, cToF(AbsoluteZeroC))
	fmt.Printf("Water freezing point: %v or %v\n", FreezingC, cToF(FreezingC))
	fmt.Printf("Water boiling point: %v or %v\n\n", BoilingC, cToF(BoilingC))

	fmt.Println("Converting ºFahrenheit to ºCelsius")
	fmt.Printf("Absolute zero: %v or %v\n", AbsoluteZeroF, fToC(AbsoluteZeroF))
	fmt.Printf("Water freezing point: %v or %v\n", FreezingF, fToC(FreezingF))
	fmt.Printf("Water boiling point: %v or %v\n", BoilingF, fToC(BoilingF))
}

func cToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func fToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%gºC", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%gºF", f)
}
