package lengthconv

// MToFt converts a meters length to feet.
func MToFt(m Meter) Feet {
	return Feet(m * 3.28084)
}

// FtToM converts a feet length to meters.
func FtToM(ft Feet) Meter {
	return Meter(ft / 3.28084)
}
