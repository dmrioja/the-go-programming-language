// Package weightconv performs kilograms and pounds conversions.
package weightconv

import "fmt"

type Kilogram float64
type Pound float64

func (kg Kilogram) String() string {
	return fmt.Sprintf("%gkg", kg)
}

func (lbs Pound) String() string {
	return fmt.Sprintf("%glbs", lbs)
}
