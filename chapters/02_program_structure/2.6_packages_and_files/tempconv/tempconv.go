// Package tempconv performs Celsius, Fahrenheit and Kelvin conversions.
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC  Celsius = -273.15
	WaterFreezingC Celsius = 0
	WaterBoilingC  Celsius = 100

	AbsoluteZeroF  Fahrenheit = -459.67
	WaterFreezingF Fahrenheit = 32
	WaterBoilingF  Fahrenheit = 212

	AbsoluteZeroK  Kelvin = 0
	WaterFreezingK Kelvin = 273.15
	WaterBoilingK  Kelvin = 373.15
)

func (c Celsius) String() string {
	return fmt.Sprintf("%gºC", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%gºF", f)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%gK", k)
}
